package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	const sourceURL = "https://raw.githubusercontent.com/haliaeetus/iso-639/master/data/iso_639-2.csv"
	inputData, err := httpGet(sourceURL)
	if err != nil {
		panic(err)
	}

	inputReader := bytes.NewReader(inputData)
	records, err := csv.NewReader(inputReader).ReadAll()
	if err != nil {
		panic(err)
	}

	languageRecords := make(languageRecordList, 0, len(records)-1)
	for _, record := range records[1:] {
		languageRecords = append(languageRecords, languageRecord{
			Name:    record[3],
			Ref:     strings.ToUpper(record[0]),
			Alpha3:  record[0],
			Alpha3B: record[1],
			Alpha2:  record[2],
		})
	}
	sort.Stable(languageRecords)
	// the source data contains some duplicate entries
	languageRecords = languageRecords.Deduplicate()

	if err = writeASTFile("code.gen.go", languageRecords.GenerateAST()); err != nil {
		panic(err)
	}
}

func writeASTFile(filename string, astFile *ast.File) (returnedErr error) {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := fd.Close(); err != nil && returnedErr == nil {
			returnedErr = err
		}
	}()
	if err := format.Node(fd, token.NewFileSet(), astFile); err != nil {
		return err
	}
	return nil
}

func httpGet(urlStr string) ([]byte, error) {
	res, err := http.Get(urlStr) // nolint: gosec
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if closeErr := res.Body.Close(); err != nil {
		return nil, closeErr
	}
	return body, err
}

type languageRecord struct {
	Name    string
	Ref     string
	Alpha3  string
	Alpha3B string
	Alpha2  string
}

func (lr languageRecord) Alpha3Ref() ast.Expr {
	if lr.Alpha3 == "" {
		return &ast.Ident{Name: empty3Name}
	}
	return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", lr.Alpha3)}
}

func (lr languageRecord) Alpha3BRef() ast.Expr {
	if lr.Alpha3B == "" {
		return &ast.Ident{Name: empty3Name}
	}
	return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", lr.Alpha3B)}
}

func (lr languageRecord) Alpha2Ref() ast.Expr {
	if lr.Alpha2 == "" {
		return &ast.Ident{Name: empty2Name}
	}
	return &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", lr.Alpha2)}
}

type languageRecordList []languageRecord

func (l languageRecordList) Len() int {
	return len(l)
}

func (l languageRecordList) Less(i, j int) bool {
	return l[i].Alpha3 < l[j].Alpha3
}

func (l languageRecordList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l languageRecordList) Deduplicate() languageRecordList {
	offset := 0
	for i := 1; i < len(l); i++ {
		if l[i-offset-1] == l[i] {
			offset++
			continue
		}
		l[i-offset] = l[i]
	}
	return l[:len(l)-offset]
}

func (l languageRecordList) GenerateAST() *ast.File {
	return &ast.File{
		Name: &ast.Ident{Name: pkgName},
		Decls: []ast.Decl{
			l.generateLanguages(),
			l.generateCodes(),
		},
	}
}

func (l languageRecordList) generateLanguages() ast.Decl {
	specs := make([]ast.Spec, 0, len(l))
	for i, lr := range l {
		specs = append(specs, &ast.ValueSpec{
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{Text: fmt.Sprintf("\n// %s is %s.", lr.Ref, lr.Name)},
				},
			},
			Names: []*ast.Ident{{Name: lr.Ref}},
			Values: []ast.Expr{
				&ast.CompositeLit{
					Type: &ast.Ident{Name: languageType},
					Elts: []ast.Expr{
						&ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", i+1)},
					},
				},
			},
		})
	}
	return &ast.GenDecl{
		Tok: token.VAR,
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "\n// Code generated. DO NOT EDIT."},
			},
		},
		Specs: specs,
	}
}

func (l languageRecordList) generateCodes() ast.Decl {
	elts := make([]ast.Expr, 0, len(l)+1)

	elts = append(elts, &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.Ident{Name: empty3Name},
			&ast.Ident{Name: empty3Name},
			&ast.Ident{Name: empty2Name},
		},
	})

	for _, lr := range l {
		elts = append(elts, &ast.CompositeLit{
			Elts: []ast.Expr{
				lr.Alpha3Ref(),
				lr.Alpha3BRef(),
				lr.Alpha2Ref(),
			},
		})
	}

	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{{Name: codes}},
				Values: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{
							Len: &ast.Ellipsis{},
							Elt: &ast.ArrayType{
								Len: &ast.Ident{Name: formatsCount},
								Elt: &ast.Ident{Name: stringType},
							},
						},
						Elts: elts,
					},
				},
			},
		},
	}
}

const pkgName = "languagecode"
const languageType = "Language"
const codes = "codes"
const formatsCount = "formatsCount"
const stringType = "string"
const empty3Name = "empty3"
const empty2Name = "empty2"
