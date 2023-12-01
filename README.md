# repro-bazel-rules_go-nogo-nilaway

We expect to be able to use //nolint comments to silence particular lint issues.

This doesn't appear to work across dependencies in some situations.

Example output:

```console
% bazel build ...
INFO: Invocation ID: 5d411c00-d1a1-40c2-893c-ffe4abc95570
INFO: Analyzed 4 targets (0 packages loaded, 0 targets configured).
INFO: Found 4 targets...
ERROR: /Volumes/Source/github.com/illicitonion/repro-bazel-rules_go-nogo-nilaway/BUILD.bazel:10:11: GoCompilePkg main_lib.a failed: (Exit 1): builder failed: error executing command (from target //:main_lib) bazel-out/darwin_arm64-opt-exec-2B5CBBC6/bin/external/go_sdk/builder_reset/builder compilepkg -sdk external/go_sdk -installsuffix darwin_arm64 -src main.go -embedroot '' -embedroot ... (remaining 23 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: nogo: errors found by nogo during build-time code analysis:
my-lib/panicboi.go:9:1: error: Potential nil panic detected. Observed nil flow from source to dereference point: 
	-> my-lib/panicboi.go:14:12: literal `nil` returned from `MakeANilPanicStruct()` in position 0
	-> __main__/main.go:12:2: result 0 of `MakeANilPanicStruct()` used as receiver to call `OhnoIMayPanic()`
	-> my-lib/panicboi.go:9:10: read by method receiver `p` accessed field `anField`
 (nilaway)
ERROR: /Volumes/Source/github.com/illicitonion/repro-bazel-rules_go-nogo-nilaway/BUILD.bazel:18:10: GoCompilePkg bin.a failed: (Exit 1): builder failed: error executing command (from target //:bin) bazel-out/darwin_arm64-opt-exec-2B5CBBC6/bin/external/go_sdk/builder_reset/builder compilepkg -sdk external/go_sdk -installsuffix darwin_arm64 -src main.go -embedroot '' -embedroot ... (remaining 23 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: nogo: errors found by nogo during build-time code analysis:
my-lib/panicboi.go:9:1: error: Potential nil panic detected. Observed nil flow from source to dereference point: 
	-> my-lib/panicboi.go:14:12: literal `nil` returned from `MakeANilPanicStruct()` in position 0
	-> __main__/main.go:12:2: result 0 of `MakeANilPanicStruct()` used as receiver to call `OhnoIMayPanic()`
	-> my-lib/panicboi.go:9:10: read by method receiver `p` accessed field `anField`
 (nilaway)
INFO: Elapsed time: 0.172s, Critical Path: 0.06s
INFO: 3 processes: 3 internal.
FAILED: Build did NOT complete successfully
```
