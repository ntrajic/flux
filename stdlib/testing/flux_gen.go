// DO NOT EDIT: This file is autogenerated via the builtin command.

package testing

import (
	flux "github.com/influxdata/flux"
	ast "github.com/influxdata/flux/ast"
)

func init() {
	flux.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 2,
					Line:   13,
				},
				File:   "testing.flux",
				Source: "package testing\n\nimport c \"csv\"\n\nbuiltin assertEquals\n\nloadStorage = (csv) => c.from(csv: csv)\nloadMem = (csv) => c.from(csv: csv)\n\ntest = (name, input, want, testFn) => {\n    got = input |> testFn()\n    return assertEquals(name: name, want: want, got: got)\n}",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 21,
						Line:   5,
					},
					File:   "testing.flux",
					Source: "builtin assertEquals",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 21,
							Line:   5,
						},
						File:   "testing.flux",
						Source: "assertEquals",
						Start: ast.Position{
							Column: 9,
							Line:   5,
						},
					},
				},
				Name: "assertEquals",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 40,
						Line:   7,
					},
					File:   "testing.flux",
					Source: "loadStorage = (csv) => c.from(csv: csv)",
					Start: ast.Position{
						Column: 1,
						Line:   7,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 12,
							Line:   7,
						},
						File:   "testing.flux",
						Source: "loadStorage",
						Start: ast.Position{
							Column: 1,
							Line:   7,
						},
					},
				},
				Name: "loadStorage",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 40,
							Line:   7,
						},
						File:   "testing.flux",
						Source: "(csv) => c.from(csv: csv)",
						Start: ast.Position{
							Column: 15,
							Line:   7,
						},
					},
				},
				Body: &ast.CallExpression{
					Arguments: []ast.Expression{&ast.ObjectExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 39,
									Line:   7,
								},
								File:   "testing.flux",
								Source: "csv: csv",
								Start: ast.Position{
									Column: 31,
									Line:   7,
								},
							},
						},
						Properties: []*ast.Property{&ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 39,
										Line:   7,
									},
									File:   "testing.flux",
									Source: "csv: csv",
									Start: ast.Position{
										Column: 31,
										Line:   7,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 34,
											Line:   7,
										},
										File:   "testing.flux",
										Source: "csv",
										Start: ast.Position{
											Column: 31,
											Line:   7,
										},
									},
								},
								Name: "csv",
							},
							Value: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 39,
											Line:   7,
										},
										File:   "testing.flux",
										Source: "csv",
										Start: ast.Position{
											Column: 36,
											Line:   7,
										},
									},
								},
								Name: "csv",
							},
						}},
					}},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 40,
								Line:   7,
							},
							File:   "testing.flux",
							Source: "c.from(csv: csv)",
							Start: ast.Position{
								Column: 24,
								Line:   7,
							},
						},
					},
					Callee: &ast.MemberExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   7,
								},
								File:   "testing.flux",
								Source: "c.from",
								Start: ast.Position{
									Column: 24,
									Line:   7,
								},
							},
						},
						Object: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   7,
									},
									File:   "testing.flux",
									Source: "c",
									Start: ast.Position{
										Column: 24,
										Line:   7,
									},
								},
							},
							Name: "c",
						},
						Property: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   7,
									},
									File:   "testing.flux",
									Source: "from",
									Start: ast.Position{
										Column: 26,
										Line:   7,
									},
								},
							},
							Name: "from",
						},
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 19,
								Line:   7,
							},
							File:   "testing.flux",
							Source: "csv",
							Start: ast.Position{
								Column: 16,
								Line:   7,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 19,
									Line:   7,
								},
								File:   "testing.flux",
								Source: "csv",
								Start: ast.Position{
									Column: 16,
									Line:   7,
								},
							},
						},
						Name: "csv",
					},
					Value: nil,
				}},
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 36,
						Line:   8,
					},
					File:   "testing.flux",
					Source: "loadMem = (csv) => c.from(csv: csv)",
					Start: ast.Position{
						Column: 1,
						Line:   8,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 8,
							Line:   8,
						},
						File:   "testing.flux",
						Source: "loadMem",
						Start: ast.Position{
							Column: 1,
							Line:   8,
						},
					},
				},
				Name: "loadMem",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 36,
							Line:   8,
						},
						File:   "testing.flux",
						Source: "(csv) => c.from(csv: csv)",
						Start: ast.Position{
							Column: 11,
							Line:   8,
						},
					},
				},
				Body: &ast.CallExpression{
					Arguments: []ast.Expression{&ast.ObjectExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 35,
									Line:   8,
								},
								File:   "testing.flux",
								Source: "csv: csv",
								Start: ast.Position{
									Column: 27,
									Line:   8,
								},
							},
						},
						Properties: []*ast.Property{&ast.Property{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 35,
										Line:   8,
									},
									File:   "testing.flux",
									Source: "csv: csv",
									Start: ast.Position{
										Column: 27,
										Line:   8,
									},
								},
							},
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 30,
											Line:   8,
										},
										File:   "testing.flux",
										Source: "csv",
										Start: ast.Position{
											Column: 27,
											Line:   8,
										},
									},
								},
								Name: "csv",
							},
							Value: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 35,
											Line:   8,
										},
										File:   "testing.flux",
										Source: "csv",
										Start: ast.Position{
											Column: 32,
											Line:   8,
										},
									},
								},
								Name: "csv",
							},
						}},
					}},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 36,
								Line:   8,
							},
							File:   "testing.flux",
							Source: "c.from(csv: csv)",
							Start: ast.Position{
								Column: 20,
								Line:   8,
							},
						},
					},
					Callee: &ast.MemberExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 26,
									Line:   8,
								},
								File:   "testing.flux",
								Source: "c.from",
								Start: ast.Position{
									Column: 20,
									Line:   8,
								},
							},
						},
						Object: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 21,
										Line:   8,
									},
									File:   "testing.flux",
									Source: "c",
									Start: ast.Position{
										Column: 20,
										Line:   8,
									},
								},
							},
							Name: "c",
						},
						Property: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 26,
										Line:   8,
									},
									File:   "testing.flux",
									Source: "from",
									Start: ast.Position{
										Column: 22,
										Line:   8,
									},
								},
							},
							Name: "from",
						},
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 15,
								Line:   8,
							},
							File:   "testing.flux",
							Source: "csv",
							Start: ast.Position{
								Column: 12,
								Line:   8,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 15,
									Line:   8,
								},
								File:   "testing.flux",
								Source: "csv",
								Start: ast.Position{
									Column: 12,
									Line:   8,
								},
							},
						},
						Name: "csv",
					},
					Value: nil,
				}},
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   13,
					},
					File:   "testing.flux",
					Source: "test = (name, input, want, testFn) => {\n    got = input |> testFn()\n    return assertEquals(name: name, want: want, got: got)\n}",
					Start: ast.Position{
						Column: 1,
						Line:   10,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 5,
							Line:   10,
						},
						File:   "testing.flux",
						Source: "test",
						Start: ast.Position{
							Column: 1,
							Line:   10,
						},
					},
				},
				Name: "test",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   13,
						},
						File:   "testing.flux",
						Source: "(name, input, want, testFn) => {\n    got = input |> testFn()\n    return assertEquals(name: name, want: want, got: got)\n}",
						Start: ast.Position{
							Column: 8,
							Line:   10,
						},
					},
				},
				Body: &ast.Block{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 2,
								Line:   13,
							},
							File:   "testing.flux",
							Source: "{\n    got = input |> testFn()\n    return assertEquals(name: name, want: want, got: got)\n}",
							Start: ast.Position{
								Column: 39,
								Line:   10,
							},
						},
					},
					Body: []ast.Statement{&ast.VariableAssignment{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 28,
									Line:   11,
								},
								File:   "testing.flux",
								Source: "got = input |> testFn()",
								Start: ast.Position{
									Column: 5,
									Line:   11,
								},
							},
						},
						ID: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 8,
										Line:   11,
									},
									File:   "testing.flux",
									Source: "got",
									Start: ast.Position{
										Column: 5,
										Line:   11,
									},
								},
							},
							Name: "got",
						},
						Init: &ast.PipeExpression{
							Argument: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 16,
											Line:   11,
										},
										File:   "testing.flux",
										Source: "input",
										Start: ast.Position{
											Column: 11,
											Line:   11,
										},
									},
								},
								Name: "input",
							},
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 28,
										Line:   11,
									},
									File:   "testing.flux",
									Source: "input |> testFn()",
									Start: ast.Position{
										Column: 11,
										Line:   11,
									},
								},
							},
							Call: &ast.CallExpression{
								Arguments: nil,
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 28,
											Line:   11,
										},
										File:   "testing.flux",
										Source: "testFn()",
										Start: ast.Position{
											Column: 20,
											Line:   11,
										},
									},
								},
								Callee: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 26,
												Line:   11,
											},
											File:   "testing.flux",
											Source: "testFn",
											Start: ast.Position{
												Column: 20,
												Line:   11,
											},
										},
									},
									Name: "testFn",
								},
							},
						},
					}, &ast.ReturnStatement{
						Argument: &ast.CallExpression{
							Arguments: []ast.Expression{&ast.ObjectExpression{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 57,
											Line:   12,
										},
										File:   "testing.flux",
										Source: "name: name, want: want, got: got",
										Start: ast.Position{
											Column: 25,
											Line:   12,
										},
									},
								},
								Properties: []*ast.Property{&ast.Property{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 35,
												Line:   12,
											},
											File:   "testing.flux",
											Source: "name: name",
											Start: ast.Position{
												Column: 25,
												Line:   12,
											},
										},
									},
									Key: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 29,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "name",
												Start: ast.Position{
													Column: 25,
													Line:   12,
												},
											},
										},
										Name: "name",
									},
									Value: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 35,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "name",
												Start: ast.Position{
													Column: 31,
													Line:   12,
												},
											},
										},
										Name: "name",
									},
								}, &ast.Property{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 47,
												Line:   12,
											},
											File:   "testing.flux",
											Source: "want: want",
											Start: ast.Position{
												Column: 37,
												Line:   12,
											},
										},
									},
									Key: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 41,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "want",
												Start: ast.Position{
													Column: 37,
													Line:   12,
												},
											},
										},
										Name: "want",
									},
									Value: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 47,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "want",
												Start: ast.Position{
													Column: 43,
													Line:   12,
												},
											},
										},
										Name: "want",
									},
								}, &ast.Property{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 57,
												Line:   12,
											},
											File:   "testing.flux",
											Source: "got: got",
											Start: ast.Position{
												Column: 49,
												Line:   12,
											},
										},
									},
									Key: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 52,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "got",
												Start: ast.Position{
													Column: 49,
													Line:   12,
												},
											},
										},
										Name: "got",
									},
									Value: &ast.Identifier{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 57,
													Line:   12,
												},
												File:   "testing.flux",
												Source: "got",
												Start: ast.Position{
													Column: 54,
													Line:   12,
												},
											},
										},
										Name: "got",
									},
								}},
							}},
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 58,
										Line:   12,
									},
									File:   "testing.flux",
									Source: "assertEquals(name: name, want: want, got: got)",
									Start: ast.Position{
										Column: 12,
										Line:   12,
									},
								},
							},
							Callee: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 24,
											Line:   12,
										},
										File:   "testing.flux",
										Source: "assertEquals",
										Start: ast.Position{
											Column: 12,
											Line:   12,
										},
									},
								},
								Name: "assertEquals",
							},
						},
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 58,
									Line:   12,
								},
								File:   "testing.flux",
								Source: "return assertEquals(name: name, want: want, got: got)",
								Start: ast.Position{
									Column: 5,
									Line:   12,
								},
							},
						},
					}},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 13,
								Line:   10,
							},
							File:   "testing.flux",
							Source: "name",
							Start: ast.Position{
								Column: 9,
								Line:   10,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 13,
									Line:   10,
								},
								File:   "testing.flux",
								Source: "name",
								Start: ast.Position{
									Column: 9,
									Line:   10,
								},
							},
						},
						Name: "name",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 20,
								Line:   10,
							},
							File:   "testing.flux",
							Source: "input",
							Start: ast.Position{
								Column: 15,
								Line:   10,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 20,
									Line:   10,
								},
								File:   "testing.flux",
								Source: "input",
								Start: ast.Position{
									Column: 15,
									Line:   10,
								},
							},
						},
						Name: "input",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 26,
								Line:   10,
							},
							File:   "testing.flux",
							Source: "want",
							Start: ast.Position{
								Column: 22,
								Line:   10,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 26,
									Line:   10,
								},
								File:   "testing.flux",
								Source: "want",
								Start: ast.Position{
									Column: 22,
									Line:   10,
								},
							},
						},
						Name: "want",
					},
					Value: nil,
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 34,
								Line:   10,
							},
							File:   "testing.flux",
							Source: "testFn",
							Start: ast.Position{
								Column: 28,
								Line:   10,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 34,
									Line:   10,
								},
								File:   "testing.flux",
								Source: "testFn",
								Start: ast.Position{
									Column: 28,
									Line:   10,
								},
							},
						},
						Name: "testFn",
					},
					Value: nil,
				}},
			},
		}},
		Imports: []*ast.ImportDeclaration{&ast.ImportDeclaration{
			As: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 9,
							Line:   3,
						},
						File:   "testing.flux",
						Source: "c",
						Start: ast.Position{
							Column: 8,
							Line:   3,
						},
					},
				},
				Name: "c",
			},
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 15,
						Line:   3,
					},
					File:   "testing.flux",
					Source: "import c \"csv\"",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 15,
							Line:   3,
						},
						File:   "testing.flux",
						Source: "\"csv\"",
						Start: ast.Position{
							Column: 10,
							Line:   3,
						},
					},
				},
				Value: "csv",
			},
		}},
		Name: "testing.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   1,
					},
					File:   "testing.flux",
					Source: "package testing",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 16,
							Line:   1,
						},
						File:   "testing.flux",
						Source: "testing",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "testing",
			},
		},
	}},
	Package: "testing",
	Path:    "testing",
}