//Copyright(c)HashiCorp,Inc.//SPDX-License-Identifier:MPL-2.0packagereleasesjsonimport("context""crypto/sha256""encoding/hex""fmt""io""log""net/http""net/url""strings""github.com/ProtonMail/go-crypto/openpgp""github.com/hashicorp/hc-install/internal/httpclient")typeChecksumDownloaderstruct{ProductVersion*ProductVersionLoggerog.LoggerArmoredPublicKeystringBaseURLstring}typeChecksumFileMapmap[string]HashSumtypeHashSum[]byte(hsHashSum)Size()int{returnlen(hs)}(hsHashSum)String()string{returnhex.EncodeToString(hs)HashSumFromHexDigest(hexDigeststring)(HashSum,error){sumBytes,err:=hex.DecodeString(hexDigest)iferr!=nil{returnnil,err}urnHashSum(sumBytes),nil}(cd*ChecksumDownloader)DownloadAndVerifyChecksums(ctxcontext.Context)(ChecksumFileMap,error){sigFilename,err:=cd.findSigFilename(cd.ProductVersion)iferr!=nil{returnnil,err}client:=httpclient.NewHTTPClient()sigURL:=fmt.Sprintf("%s/%s/%s/%s",cd.BaseURL,url.PathEscape(cd.ProductVersion.Name),url.PathEscape(cd.ProductVersion.RawVersion),url.PathEscape(sigFilename))cd.Logger.Printf("downloadingsignaturefrom%s",sigURL)req,err:=http.NewRequestWithContext(ctx,http.MethodGet,sigURL,nil)iferr!=nil{returnnil,fmt.Errorf("failedtocreaterequestfor%q:%w",sigURL,err)}sigResp,err:=client.Do(req)iferr!=nil{returnnil,err}ifsigResp.StatusCode!=200{returnnil,fmt.Errorf("failedtodownloadsignaturefrom%q:%s",sigURL,sigResp.Status)}defersigResp.Body.Close()shasumsURL:=fmt.Sprintf("%s/%s/%s/%s",cd.BaseURL,url.PathEscape(cd.ProductVersion.Name),url.PathEscape(cd.ProductVersion.RawVersion),url.PathEscape(cd.ProductVersion.SHASUMS))cd.Logger.Printf("downloadingchecksumsfrom%s",shasumsURL)req,err=http.NewRequestWithContext(ctx,http.MethodGet,shasumsURL,nil)iferr!=nil{returnnil,fmt.Errorf("failedtocreaterequestfor%q:%w",shasumsURL,err)}sumsResp,err:=client.Do(req)iferr!=nil{returnnil,err}ifsumsResp.StatusCode!=200{returnnil,fmt.Errorf("failedtodownloadchecksumsfrom%q:%s",shasumsURL,sumsResp.Status)}defersumsResp.Body.Close()varshaSumsstrings.BuildersumsReader:=io.TeeReader(sumsResp.Body,&shaSums)err=cd.verifySumsSignature(sumsReader,sigResp.Body)iferr!=nil{returnnil,err}returnfileMapFromChecksums(shaSums)}fileMapFromChecksums(checksumsstrings.Builder)(ChecksumFileMap,error){csMap:=make(ChecksumFileMap,0)lines:=strings.Split(checksums.String(),"\n")for_,line:=rangelines{line=strings.TrimSpace(line)ifline==""{continue}parts:=strings.Fields(line)iflen(parts)!=2{returnnil,fmt.Errorf("unexpectedchecksumlineformat:%q",line)}h,err:=HashSumFromHexDigest(parts[0])iferr!=nil{returnnil,err}ifh.Size()!=sha256.Size{returnnil,fmt.Errorf("unexpectedsha256format(len:%d,expected:%d)",h.Size(),sha256.Size)}Map[parts[1]]=h}returncsMap,nil}(cd*ChecksumDownloader)verifySumsSignature(checksums,signatureio.Reader)error{el,err:=cd.keyEntityList()iferr!=nil{returnerr}_,err=openpgp.CheckDetachedSignature(el,checksums,signature,nil)iferr!=nil{returnfmt.Errorf("unabletoverifychecksumssignature:%w",err)}cd.Logger.Printf("checksumsignatureisvalid")returnnil}(cd*ChecksumDownloader)findSigFilename(pv*ProductVersion)(string,error){sigFiles:=pv.SHASUMSSigsiflen(sigFiles)==0{sigFiles=[]string{pv.SHASUMSSig}}keyIds,err:=cd.pubKeyIds()iferr!=nil{return"",err}for_,filename:=rangesigFiles{for_,keyID:=rangekeyIds{ifstrings.HasSuffix(filename,fmt.Sprintf("_SHA256SUMS.%s.sig",keyID)){returnfilename,nil}}ifstrings.HasSuffix(filename,"_SHA256SUMS.sig"){eturnfilename,nil}}return"",fmt.Errorf("nosuitablesigfilefound")}(cd*ChecksumDownloader)pubKeyIds()([]string,error){entityList,err:=cd.keyEntityList()iferr!=nil{returnnil,err}gerprints:=make([]string,0)for_,entity:=rangeentityList{fingerprints=append(fingerprints,entity.PrimaryKey.KeyIdShortString())}returnfingerprints,nil}(cd*ChecksumDownloader)keyEntityList()(openpgp.EntityList,error){ifcd.ArmoredPublicKey==""{returnnil,fmt.Errorf("nopublickeyprovided")}returnopenpgp.ReadArmoredKeyRing(strings.NewReader(cd.ArmoredPublicKey))}