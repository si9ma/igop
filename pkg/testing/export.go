// export by github.com/goplus/interp/cmd/qexp

package testing

import (
	"testing"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("testing", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*testing.B).Cleanup":                        (*testing.B).Cleanup,
	"(*testing.B).Error":                          (*testing.B).Error,
	"(*testing.B).Errorf":                         (*testing.B).Errorf,
	"(*testing.B).Fail":                           (*testing.B).Fail,
	"(*testing.B).FailNow":                        (*testing.B).FailNow,
	"(*testing.B).Failed":                         (*testing.B).Failed,
	"(*testing.B).Fatal":                          (*testing.B).Fatal,
	"(*testing.B).Fatalf":                         (*testing.B).Fatalf,
	"(*testing.B).Helper":                         (*testing.B).Helper,
	"(*testing.B).Log":                            (*testing.B).Log,
	"(*testing.B).Logf":                           (*testing.B).Logf,
	"(*testing.B).Name":                           (*testing.B).Name,
	"(*testing.B).ReportAllocs":                   (*testing.B).ReportAllocs,
	"(*testing.B).ReportMetric":                   (*testing.B).ReportMetric,
	"(*testing.B).ResetTimer":                     (*testing.B).ResetTimer,
	"(*testing.B).Run":                            (*testing.B).Run,
	"(*testing.B).RunParallel":                    (*testing.B).RunParallel,
	"(*testing.B).SetBytes":                       (*testing.B).SetBytes,
	"(*testing.B).SetParallelism":                 (*testing.B).SetParallelism,
	"(*testing.B).Skip":                           (*testing.B).Skip,
	"(*testing.B).SkipNow":                        (*testing.B).SkipNow,
	"(*testing.B).Skipf":                          (*testing.B).Skipf,
	"(*testing.B).Skipped":                        (*testing.B).Skipped,
	"(*testing.B).StartTimer":                     (*testing.B).StartTimer,
	"(*testing.B).StopTimer":                      (*testing.B).StopTimer,
	"(*testing.M).Run":                            (*testing.M).Run,
	"(*testing.PB).Next":                          (*testing.PB).Next,
	"(*testing.T).Cleanup":                        (*testing.T).Cleanup,
	"(*testing.T).Error":                          (*testing.T).Error,
	"(*testing.T).Errorf":                         (*testing.T).Errorf,
	"(*testing.T).Fail":                           (*testing.T).Fail,
	"(*testing.T).FailNow":                        (*testing.T).FailNow,
	"(*testing.T).Failed":                         (*testing.T).Failed,
	"(*testing.T).Fatal":                          (*testing.T).Fatal,
	"(*testing.T).Fatalf":                         (*testing.T).Fatalf,
	"(*testing.T).Helper":                         (*testing.T).Helper,
	"(*testing.T).Log":                            (*testing.T).Log,
	"(*testing.T).Logf":                           (*testing.T).Logf,
	"(*testing.T).Name":                           (*testing.T).Name,
	"(*testing.T).Parallel":                       (*testing.T).Parallel,
	"(*testing.T).Run":                            (*testing.T).Run,
	"(*testing.T).Skip":                           (*testing.T).Skip,
	"(*testing.T).SkipNow":                        (*testing.T).SkipNow,
	"(*testing.T).Skipf":                          (*testing.T).Skipf,
	"(*testing.T).Skipped":                        (*testing.T).Skipped,
	"(testing.BenchmarkResult).AllocedBytesPerOp": (testing.BenchmarkResult).AllocedBytesPerOp,
	"(testing.BenchmarkResult).AllocsPerOp":       (testing.BenchmarkResult).AllocsPerOp,
	"(testing.BenchmarkResult).MemString":         (testing.BenchmarkResult).MemString,
	"(testing.BenchmarkResult).NsPerOp":           (testing.BenchmarkResult).NsPerOp,
	"(testing.BenchmarkResult).String":            (testing.BenchmarkResult).String,
	"(testing.TB).Cleanup":                        (testing.TB).Cleanup,
	"(testing.TB).Error":                          (testing.TB).Error,
	"(testing.TB).Errorf":                         (testing.TB).Errorf,
	"(testing.TB).Fail":                           (testing.TB).Fail,
	"(testing.TB).FailNow":                        (testing.TB).FailNow,
	"(testing.TB).Failed":                         (testing.TB).Failed,
	"(testing.TB).Fatal":                          (testing.TB).Fatal,
	"(testing.TB).Fatalf":                         (testing.TB).Fatalf,
	"(testing.TB).Helper":                         (testing.TB).Helper,
	"(testing.TB).Log":                            (testing.TB).Log,
	"(testing.TB).Logf":                           (testing.TB).Logf,
	"(testing.TB).Name":                           (testing.TB).Name,
	"(testing.TB).Skip":                           (testing.TB).Skip,
	"(testing.TB).SkipNow":                        (testing.TB).SkipNow,
	"(testing.TB).Skipf":                          (testing.TB).Skipf,
	"(testing.TB).Skipped":                        (testing.TB).Skipped,
	"testing.AllocsPerRun":                        testing.AllocsPerRun,
	"testing.Benchmark":                           testing.Benchmark,
	"testing.CoverMode":                           testing.CoverMode,
	"testing.Coverage":                            testing.Coverage,
	"testing.Init":                                testing.Init,
	"testing.Main":                                testing.Main,
	"testing.MainStart":                           testing.MainStart,
	"testing.RegisterCover":                       testing.RegisterCover,
	"testing.RunBenchmarks":                       testing.RunBenchmarks,
	"testing.RunExamples":                         testing.RunExamples,
	"testing.RunTests":                            testing.RunTests,
	"testing.Short":                               testing.Short,
	"testing.Verbose":                             testing.Verbose,
}

var typList = []interface{}{
	(*testing.B)(nil),
	(*testing.BenchmarkResult)(nil),
	(*testing.Cover)(nil),
	(*testing.CoverBlock)(nil),
	(*testing.InternalBenchmark)(nil),
	(*testing.InternalExample)(nil),
	(*testing.InternalTest)(nil),
	(*testing.M)(nil),
	(*testing.PB)(nil),
	(*testing.T)(nil),
	(*testing.TB)(nil),
}