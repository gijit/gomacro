/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * stmt_3_test.go
 *
 *  Created on May 01, 2018
 *      Author Massimiliano Ghilardi
 */

package experiments

import (
	r "reflect"
	"testing"
)

type (
	Env7 struct {
		Binds []r.Value
		Outer *Env7
	}
	Run7 struct {
		Env       *Env7
		Code      []Stmt7
		IP        int
		Signal    int
		Interrupt Stmt7
	}
	Stmt7 func(run *Run7) Stmt7
)

func nop7(run *Run7) Stmt7 {
	run.IP++
	return run.Code[run.IP]
}

func interrupt7(run *Run7) Stmt7 {
	run.Signal = 1
	return run.Interrupt
}

func newRun7() *Run7 {
	env := &Env7{
		Binds: make([]r.Value, 10),
	}
	code := make([]Stmt7, n+1)
	for i := 0; i < n; i++ {
		code[i] = nop7
	}
	code[n] = nil
	return &Run7{Env: env, Code: code}
}

func BenchmarkStmt7(b *testing.B) {
	run := newRun7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		run.IP = 0
		stmt := run.Code[0]
		for {
			if stmt = stmt(run); stmt == nil {
				break
			}
		}
	}
}

func BenchmarkStmt7Unroll(b *testing.B) {
	run := newRun7()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		run.IP = 0
		stmt := run.Code[0]
		for {
			if stmt = stmt(run); stmt != nil {
				if stmt = stmt(run); stmt != nil {
					if stmt = stmt(run); stmt != nil {
						if stmt = stmt(run); stmt != nil {
							if stmt = stmt(run); stmt != nil {
								if stmt = stmt(run); stmt != nil {
									if stmt = stmt(run); stmt != nil {
										if stmt = stmt(run); stmt != nil {
											if stmt = stmt(run); stmt != nil {
												if stmt = stmt(run); stmt != nil {
													if stmt = stmt(run); stmt != nil {
														if stmt = stmt(run); stmt != nil {
															if stmt = stmt(run); stmt != nil {
																continue
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			break
		}
	}
}

func BenchmarkStmt7Spin(b *testing.B) {
	run := newRun7()

	run.Code[n] = interrupt7
	run.Interrupt = interrupt7

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		run.IP = 0
		run.Signal = 0
		stmt := run.Code[0]
		for {
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			if run.Signal != 0 {
				break
			}
		}
	}
}

func BenchmarkStmt7Adaptive13(b *testing.B) {
	run := newRun7()

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		run.IP = 0
		run.Signal = 0
		stmt := run.Code[0]
		if stmt == nil {
			continue outer
		}
		for j := 0; j < 5; j++ {
			if stmt = stmt(run); stmt != nil {
				if stmt = stmt(run); stmt != nil {
					if stmt = stmt(run); stmt != nil {
						if stmt = stmt(run); stmt != nil {
							if stmt = stmt(run); stmt != nil {
								if stmt = stmt(run); stmt != nil {
									if stmt = stmt(run); stmt != nil {
										if stmt = stmt(run); stmt != nil {
											if stmt = stmt(run); stmt != nil {
												if stmt = stmt(run); stmt != nil {
													if stmt = stmt(run); stmt != nil {
														if stmt = stmt(run); stmt != nil {
															if stmt = stmt(run); stmt != nil {
																continue
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			continue outer
		}
		run.Code[n] = interrupt7
		run.Interrupt = interrupt7
		for {
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)

			if run.Signal != 0 {
				continue outer
			}
		}
	}
}

func BenchmarkStmt7Adaptive19(b *testing.B) {
	run := newRun7()

	b.ResetTimer()
outer:
	for i := 0; i < b.N; i++ {
		run.IP = 0
		run.Signal = 0
		stmt := run.Code[0]
		if stmt == nil {
			continue outer
		}
		for j := 0; j < 5; j++ {
			if stmt = stmt(run); stmt != nil {
				if stmt = stmt(run); stmt != nil {
					if stmt = stmt(run); stmt != nil {
						if stmt = stmt(run); stmt != nil {
							if stmt = stmt(run); stmt != nil {
								if stmt = stmt(run); stmt != nil {
									if stmt = stmt(run); stmt != nil {
										if stmt = stmt(run); stmt != nil {
											if stmt = stmt(run); stmt != nil {
												if stmt = stmt(run); stmt != nil {
													if stmt = stmt(run); stmt != nil {
														if stmt = stmt(run); stmt != nil {
															if stmt = stmt(run); stmt != nil {
																if stmt = stmt(run); stmt != nil {
																	if stmt = stmt(run); stmt != nil {
																		if stmt = stmt(run); stmt != nil {
																			if stmt = stmt(run); stmt != nil {
																				if stmt = stmt(run); stmt != nil {
																					if stmt = stmt(run); stmt != nil {
																						continue
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			continue outer
		}
		run.Code[n] = interrupt7
		run.Interrupt = interrupt7
		for {
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)
			stmt = stmt(run)

			if run.Signal != 0 {
				continue outer
			}
		}
	}
}
