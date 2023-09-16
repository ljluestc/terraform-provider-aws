// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0//go:build generate// +build generatepackage mainimport (	_ "embed"	"fmt"	"sort"	"strings"	"github.com/hashicorp/terraform-provider-aws/internal/generate/common")//go:embed header.tmplvar header string//go:embed file.tmplvar tmpl stringconst (maxBadCaps = 31)type CapsDatum struct {Wrong stringRight stringTest  string}type TemplateData struct {BadCaps []CapsDatum}


func main() {const (filename     = "../../../names/caps.md"capsDataFile = "../../../names/caps.csv")g := common.NewGenerator()g.Infof("Generating %s", strings.TrimPrefix(filename, "../../../"))badCaps, err := readBadCaps(capsDataFile)if err != nil {g.Fatalf("error reading %s: %s", capsDataFile, err)}td := TemplateData{}td.BadCaps = badCapsd := g.NewUnformattedFileDestination(filename)if err := d.WriteTemplate("namescapslist", header+"\n"+tmpl+"\n", td); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}if err := d.Write(); err != nil {g.Fatalf("generating file (%s): %s", filename, err)}}


func readBadCaps(filename string) ([]CapsDatum, error) {caps, err := common.ReadAllCSVData(filename)if err != nil {return nil, err}var capsList []CapsDatumfor i, row := range caps {if i < 1 { // skip headercontinue}// 0 - wrong// 1 - rightif row[0] == "" {continue}capsList = append(capsList, CapsDatum{Wrong: row[0],Right: row[1],})}sort.SliceStable(capsList, 


func(i, j int) bool {if len(capsList[i].Wrong) == len(capsList[j].Wrong) {return capsList[i].Wrong < capsList[j].Wrong}return len(capsList[j].Wrong) < len(capsList[i].Wrong)})onChunk := -1for i := range capsList {if i%maxBadCaps == 0 {onChunk++}capsList[i].Test = fmt.Sprintf(`%s%d`, "caps", onChunk)}sort.SliceStable(capsList, 


func(i, j int) bool {if strings.EqualFold(capsList[i].Wrong, capsList[j].Wrong) {return capsList[i].Wrong < capsList[j].Wrong}return strings.ToLower(capsList[i].Wrong) < strings.ToLower(capsList[j].Wrong)})return capsList, nil}