// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0//go:build generate// +build generatepackage mainimport (	"flag"	"fmt"	"os"	"strings"	"time"	"github.com/YakDriver/regexache"	"github.com/hashicorp/terraform-provider-aws/internal/generate/common"	v1 "github.com/hashicorp/terraform-provider-aws/internal/generate/tags/templates/v1"	v2 "github.com/hashicorp/terraform-provider-aws/internal/generate/tags/templates/v2"	"github.com/hashicorp/terraform-provider-aws/names")const (sdkV1 = 1sdkV2 = 2)const (defaultListTags
func           = "listTags"defaultUpdateTags
func         = "updateTags"defaultWaitTagsPropagated
func = "waitTagsPropagated")var (createTags               = flag.Bool("CreateTags", false, "whether to generate CreateTags")getTag                   = flag.Bool("GetTag", false, "whether to generate GetTag")listTags                 = flag.Bool("ListTags", false, "whether to generate ListTags")serviceTagsMap           = flag.Bool("ServiceTagsMap", false, "whether to generate service tags for map")serviceTagsSlice         = flag.Bool("ServiceTagsSlice", false, "whether to generate service tags for slice")untagInNeedTagType       = flag.Bool("UntagInNeedTagType", false, "whether Untag input needs tag type")updateTags               = flag.Bool("UpdateTags", false, "whether to generate UpdateTags")updateTagsNoIgnoreSystem = flag.Bool("UpdateTagsNoIgnoreSystem", false, "whether to not ignore system tags in UpdateTags")waitForPropagation       = flag.Bool("Wait", false, "whether to generate WaitTagsPropagated")createTags
func          = flag.String("CreateTags
func", "createTags", "createTags
func")getTag
func              = flag.String("GetTag
func", "GetTag", "getTag
func")getTagsIn
func           = flag.String("GetTagsIn
func", "getTagsIn", "getTagsIn
func")keyValueTags
func        = flag.String("KeyValueTags
func", "KeyValueTags", "keyValueTags
func")listTags
func            = flag.String("ListTags
func", defaultListTags
func, "listTags
func")listTagsInFiltIDName    = flag.String("ListTagsInFiltIDName", "", "listTagsInFiltIDName")listTagsInIDElem        = flag.String("ListTagsInIDElem", "ResourceArn", "listTagsInIDElem")listTagsInIDNeedSlice   = flag.String("ListTagsInIDNeedSlice", "", "listTagsInIDNeedSlice")listTagsOp              = flag.String("ListTagsOp", "ListTagsForResource", "listTagsOp")listTagsOpPaginated     = flag.Bool("ListTagsOpPaginated", false, "whether ListTagsOp is paginated")listTagsOutTagsElem     = flag.String("ListTagsOutTagsElem", "Tags", "listTagsOutTagsElem")setTagsOut
func          = flag.String("SetTagsOut
func", "setTagsOut", "setTagsOut
func")tagInCustomVal          = flag.String("TagInCustomVal", "", "tagInCustomVal")tagInIDElem             = flag.String("TagInIDElem", "ResourceArn", "tagInIDElem")tagInIDNeedSlice        = flag.String("TagInIDNeedSlice", "", "tagInIDNeedSlice")tagInIDNeedValueSlice   = flag.String("TagInIDNeedValueSlice", "", "tagInIDNeedValueSlice")tagInTagsElem           = flag.String("TagInTagsElem", "Tags", "tagInTagsElem")tagKeyType              = flag.String("TagKeyType", "", "tagKeyType")tagOp                   = flag.String("TagOp", "TagResource", "tagOp")tagOpBatchSize          = flag.String("TagOpBatchSize", "", "tagOpBatchSize")tagResTypeElem          = flag.String("TagResTypeElem", "", "tagResTypeElem")tagType                 = flag.String("TagType", "Tag", "tagType")tagType2                = flag.String("TagType2", "", "tagType")tagTypeAddBoolElem      = flag.String("TagTypeAddBoolElem", "", "TagTypeAddBoolElem")tagTypeIDElem           = flag.String("TagTypeIDElem", "", "tagTypeIDElem")tagTypeKeyElem          = flag.String("TagTypeKeyElem", "Key", "tagTypeKeyElem")tagTypeValElem          = flag.String("TagTypeValElem", "Value", "tagTypeValElem")tags
func                = flag.String("Tags
func", "Tags", "tags
func")untagInCustomVal        = flag.String("UntagInCustomVal", "", "untagInCustomVal")untagInNeedTagKeyType   = flag.String("UntagInNeedTagKeyType", "", "untagInNeedTagKeyType")untagInTagsElem         = flag.String("UntagInTagsElem", "TagKeys", "untagInTagsElem")untagOp                 = flag.String("UntagOp", "UntagResource", "untagOp")updateTags
func          = flag.String("UpdateTags
func", defaultUpdateTags
func, "updateTags
func")waitTagsPropagated
func  = flag.String("Wait
func", defaultWaitTagsPropagated
func, "wait
func")waitContinuousOccurence = flag.Int("WaitContinuousOccurence", 0, "ContinuousTargetOccurence for Wait 
function")waitDelay               = flag.Duration("WaitDelay", 0, "Delay for Wait 
function")waitMinTimeout          = flag.Duration("WaitMinTimeout", 0, `"MinTimeout" (minimum poll interval) for Wait 
function`)waitPollInterval        = flag.Duration("WaitPollInterval", 0, "PollInterval for Wait 
function")waitTimeout             = flag.Duration("WaitTimeout", 0, "Timeout for Wait 
function")parentNotFoundErrCode = flag.String("ParentNotFoundErrCode", "", "Parent 'NotFound' Error Code")parentNotFoundErrMsg  = flag.String("ParentNotFoundErrMsg", "", "Parent 'NotFound' Error Message")sdkServicePackage = flag.String("AWSSDKServicePackage", "", "AWS Go SDK package to use. Defaults to the provider service package name.")sdkVersion        = flag.Int("AWSSDKVersion", sdkV1, "Version of the AWS Go SDK to use i.e. 1 or 2")kvtValues         = flag.Bool("KVTValues", false, "Whether KVT string map is of string pointers")skipAWSImp        = flag.Bool("SkipAWSImp", false, "Whether to skip importing the AWS Go SDK aws package") // nosemgrep:ci.aws-in-var-nameskipNamesImp      = flag.Bool("SkipNamesImp", false, "Whether to skip importing names")skipServiceImp    = flag.Bool("SkipAWSServiceImp", false, "Whether to skip importing the AWS service package")skipTypesImp      = flag.Bool("SkipTypesImp", false, "Whether to skip importing types"))
func usage() {fmt.Fprintf(os.Stderr, "Usage:\n")fmt.Fprintf(os.Stderr, "\tmain.go [flags]\n\n")fmt.Fprintf(os.Stderr, "Flags:\n")flag.PrintDefaults()}type TemplateBody struct {getTag             stringheader             stringlistTags           stringserviceTagsMap     stringserviceTagsSlice   stringupdateTags         stringwaitTagsPropagated string}
func newTemplateBody(version int, kvtValues bool) *TemplateBody {switch version {case sdkV1:return &TemplateBody{"\n" + v1.GetTagBody,v1.HeaderBody,"\n" + v1.ListTagsBody,"\n" + v1.ServiceTagsMapBody,"\n" + v1.ServiceTagsSliceBody,"\n" + v1.UpdateTagsBody,"\n" + v1.WaitTagsPropagatedBody,}case sdkV2:if kvtValues {return &TemplateBody{"\n" + v2.GetTagBody,v2.HeaderBody,"\n" + v2.ListTagsBody,"\n" + v2.ServiceTagsValueMapBody,"\n" + v2.ServiceTagsSliceBody,"\n" + v2.UpdateTagsBody,"\n" + v2.WaitTagsPropagatedBody,}}return &TemplateBody{"\n" + v2.GetTagBody,v2.HeaderBody,"\n" + v2.ListTagsBody,"\n" + v2.ServiceTagsMapBody,"\n" + v2.ServiceTagsSliceBody,"\n" + v2.UpdateTagsBody,"\n" + v2.WaitTagsPropagatedBody,}default:return nil}}type TemplateData struct {AWSService             stringAWSServiceIfacePackage stringClientType             stringProviderNameUpper      stringServicePackage         stringCreateTags
func          stringGetTag
func              stringGetTagsIn
func           stringKeyValueTags
func        stringListTags
func            stringListTagsInFiltIDName    stringListTagsInIDElem        stringListTagsInIDNeedSlice   stringListTagsOp              stringListTagsOpPaginated     boolListTagsOutTagsElem     stringParentNotFoundErrCode   stringParentNotFoundErrMsg    stringRetryCreateOnNotFound   stringSetTagsOut
func          stringTagInCustomVal          stringTagInIDElem             stringTagInIDNeedSlice        stringTagInIDNeedValueSlice   stringTagInTagsElem           stringTagKeyType              stringTagOp                   stringTagOpBatchSize          stringTagPackage              stringTagResTypeElem          stringTagType                 stringTagType2                stringTagTypeAddBoolElem      stringTagTypeAddBoolElemSnake stringTagTypeIDElem           stringTagTypeKeyElem          stringTagTypeValElem          stringTags
func                stringUntagInCustomVal        stringUntagInNeedTagKeyType   stringUntagInNeedTagType      boolUntagInTagsElem         stringUntagOp                 stringUpdateTags
func          stringUpdateTagsIgnoreSystem  boolWaitForPropagation      boolWaitTagsPropagated
func  stringWaitContinuousOccurence intWaitDelay               stringWaitMinTimeout          stringWaitPollInterval        stringWaitTimeout             string// The following are specific to writing import paths in the `headerBody`;// to include the package, set the corresponding field's value to trueConnsPkg         boolFmtPkg           boolHelperSchemaPkg  boolInternalTypesPkg boolLoggingPkg       boolNamesPkg         boolSkipAWSImp       boolSkipServiceImp   boolSkipTypesImp     boolTfLogPkg         boolTfResourcePkg    boolTimePkg          boolIsDefaultListTags   boolIsDefaultUpdateTags bool}
func main() {flag.Usage = usageflag.Parse()filename := `tags_gen.go`if args := flag.Args(); len(args) > 0 {filename = args[0]}g := common.NewGenerator()if *sdkVersion != sdkV1 && *sdkVersion != sdkV2 {g.Fatalf("AWS SDK Go Version %d not supported", *sdkVersion)}servicePackage := os.Getenv("GOPACKAGE")if *sdkServicePackage == "" {sdkServicePackage = &servicePackage}g.Infof("Generating internal/service/%s/%s", servicePackage, filename)awsPkg, err := names.AWSGoPackage(*sdkServicePackage, *sdkVersion)if err != nil {g.Fatalf("encountered: %s", err)}var awsIntfPkg stringif *sdkVersion == sdkV1 && (*getTag || *listTags || *updateTags) {awsIntfPkg = fmt.Sprintf("%[1]s/%[1]siface", awsPkg)}clientTypeName, err := names.AWSGoClientTypeName(*sdkServicePackage, *sdkVersion)if err != nil {g.Fatalf("encountered: %s", err)}providerNameUpper, err := names.ProviderNameUpper(*sdkServicePackage)if err != nil {g.Fatalf("encountered: %s", err)}createTags
func := *createTags
funcif *createTags && !*updateTags {g.Infof("CreateTags only valid with UpdateTags")createTags
func = ""} else if !*createTags {createTags
func = ""}var clientType stringif *sdkVersion == sdkV1 {clientType = fmt.Sprintf("%siface.%sAPI", awsPkg, clientTypeName)} else {clientType = fmt.Sprintf("*%s.%s", awsPkg, clientTypeName)}tagPackage := awsPkgif tagPackage == "wafregional" {tagPackage = "waf"if *sdkVersion == sdkV1 {awsPkg = ""}}templateData := TemplateData{AWSService:             awsPkg,AWSServiceIfacePackage: awsIntfPkg,ClientType:             clientType,ProviderNameUpper:      providerNameUpper,ServicePackage:         servicePackage,ConnsPkg:         (*listTags && *listTags
func == defaultListTags
func) || (*updateTags && *updateTags
func == defaultUpdateTags
func),FmtPkg:           *updateTags,HelperSchemaPkg:  awsPkg == "autoscaling",InternalTypesPkg: (*listTags && *listTags
func == defaultListTags
func) || *serviceTagsMap || *serviceTagsSlice,LoggingPkg:       *updateTags,NamesPkg:         *updateTags && !*skipNamesImp,SkipAWSImp:       *skipAWSImp,SkipServiceImp:   *skipServiceImp,SkipTypesImp:     *skipTypesImp,TfLogPkg:         *updateTags,TfResourcePkg:    (*getTag || *waitForPropagation),TimePkg:          *waitForPropagation,CreateTags
func:          createTags
func,GetTag
func:              *getTag
func,GetTagsIn
func:           *getTagsIn
func,KeyValueTags
func:        *keyValueTags
func,ListTags
func:            *listTags
func,ListTagsInFiltIDName:    *listTagsInFiltIDName,ListTagsInIDElem:        *listTagsInIDElem,ListTagsInIDNeedSlice:   *listTagsInIDNeedSlice,ListTagsOp:              *listTagsOp,ListTagsOpPaginated:     *listTagsOpPaginated,ListTagsOutTagsElem:     *listTagsOutTagsElem,ParentNotFoundErrCode:   *parentNotFoundErrCode,ParentNotFoundErrMsg:    *parentNotFoundErrMsg,SetTagsOut
func:          *setTagsOut
func,TagInCustomVal:          *tagInCustomVal,TagInIDElem:             *tagInIDElem,TagInIDNeedSlice:        *tagInIDNeedSlice,TagInIDNeedValueSlice:   *tagInIDNeedValueSlice,TagInTagsElem:           *tagInTagsElem,TagKeyType:              *tagKeyType,TagOp:                   *tagOp,TagOpBatchSize:          *tagOpBatchSize,TagPackage:              tagPackage,TagResTypeElem:          *tagResTypeElem,TagType:                 *tagType,TagType2:                *tagType2,TagTypeAddBoolElem:      *tagTypeAddBoolElem,TagTypeAddBoolElemSnake: toSnakeCase(*tagTypeAddBoolElem),TagTypeIDElem:           *tagTypeIDElem,TagTypeKeyElem:          *tagTypeKeyElem,TagTypeValElem:          *tagTypeValElem,Tags
func:                *tags
func,UntagInCustomVal:        *untagInCustomVal,UntagInNeedTagKeyType:   *untagInNeedTagKeyType,UntagInNeedTagType:      *untagInNeedTagType,UntagInTagsElem:         *untagInTagsElem,UntagOp:                 *untagOp,UpdateTags
func:          *updateTags
func,UpdateTagsIgnoreSystem:  !*updateTagsNoIgnoreSystem,WaitForPropagation:      *waitForPropagation,WaitTagsPropagated
func:  *waitTagsPropagated
func,WaitContinuousOccurence: *waitContinuousOccurence,WaitDelay:               formatDuration(*waitDelay),WaitMinTimeout:          formatDuration(*waitMinTimeout),WaitPollInterval:        formatDuration(*waitPollInterval),WaitTimeout:             formatDuration(*waitTimeout),IsDefaultListTags:   *listTags
func == defaultListTags
func,IsDefaultUpdateTags: *updateTags
func == defaultUpdateTags
func,}templateBody := newTemplateBody(*sdkVersion, *kvtValues)d := g.NewGoFileDestination(filename)if *getTag || *listTags || *serviceTagsMap || *serviceTagsSlice || *updateTags {// If you intend to only generate Tags and KeyValueTags helper methods,// the corresponding aws-sdk-goservice package does not need to be importedif !*getTag && !*listTags && !*serviceTagsSlice && !*updateTags {templateData.AWSService = ""templateData.TagPackage = ""}if err := d.WriteTemplate("header", templateBody.header, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *getTag {if err := d.WriteTemplate("gettag", templateBody.getTag, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *listTags {if err := d.WriteTemplate("listtags", templateBody.listTags, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *serviceTagsMap {if err := d.WriteTemplate("servicetagsmap", templateBody.serviceTagsMap, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *serviceTagsSlice {if err := d.WriteTemplate("servicetagsslice", templateBody.serviceTagsSlice, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *updateTags {if err := d.WriteTemplate("updatetags", templateBody.updateTags, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if *waitForPropagation {if err := d.WriteTemplate("waittagspropagated", templateBody.waitTagsPropagated, templateData); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}if err := d.Write(); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}
func toSnakeCase(str string) string {result := regexache.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1}_${2}")result = regexache.MustCompile("([0-9a-z])([A-Z])").ReplaceAllString(result, "${1}_${2}")return strings.ToLower(result)}
func formatDuration(d time.Duration) string {if d == 0 {return ""}var buf []stringif h := d.Hours(); h >= 1 {buf = append(buf, fmt.Sprintf("%d * time.Hour", int64(h)))d = d - time.Duration(int64(h)*int64(time.Hour))}if m := d.Minutes(); m >= 1 {buf = append(buf, fmt.Sprintf("%d * time.Minute", int64(m)))d = d - time.Duration(int64(m)*int64(time.Minute))}if s := d.Seconds(); s >= 1 {buf = append(buf, fmt.Sprintf("%d * time.Second", int64(s)))d = d - time.Duration(int64(s)*int64(time.Second))}if ms := d.Milliseconds(); ms >= 1 {buf = append(buf, fmt.Sprintf("%d * time.Millisecond", int64(ms)))}// Ignoring anything below millisecondsreturn strings.Join(buf, " + ")}