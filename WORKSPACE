load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_gazelle",
    sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d1ffd055969c8f8d431e2d439813e42326961d0942bdf734d2c95dc30c369566",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.5/rules_go-v0.24.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.5/rules_go-v0.24.5.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_remyoudompheng_bigfft",
    importpath = "github.com/remyoudompheng/bigfft",
    sum = "h1:OdAsTTz6OkFY5QxjkYwrChwuRruF69c169dPK26NUlk=",
    version = "v0.0.0-20200410134404-eec4a21b6bb0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=",
    version = "v0.0.0-20200622213623-75b288015ac9",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:uwuIcX0g4Yl1NC5XAz37xsr2lTtcqevgzYNVt49waME=",
    version = "v0.0.0-20201110031124-69a78807bb2b",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:+Nyd8tzPX9R7BWHguqsrbFdRx3WQ/1ib8I44HXV5yTA=",
    version = "v0.0.0-20200930185726-fdedc70b468f",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k=",
    version = "v0.3.3",
)

# go_repository(
#     name = "org_golang_x_tools",
#     importpath = "golang.org/x/tools",
#     sum = "h1:FDhOuMEY4JVRztM/gsbk+IKUQ8kj74bxZrgw87eMMVc=",
#     version = "v0.0.0-20180917221912-90fa682c2a6e",
# )

go_repository(
    name = "org_modernc_ccgo",
    importpath = "modernc.org/ccgo",
    sum = "h1:aIU6fp+ic9v4M6l6IAb0LD8byPDmtOhKXRnrNwkp88o=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_ccir",
    importpath = "modernc.org/ccir",
    sum = "h1:fAushdwIOmC+RLDpcFRp26UPHHJbvO4AQ5vt8BUZEyE=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_internal",
    importpath = "modernc.org/internal",
    sum = "h1:XMDsFDcBDsibbBnHB2xzljZ+B1yrOVLEFkKL2u15Glw=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_mathutil",
    importpath = "modernc.org/mathutil",
    sum = "h1:FeylZSVX8S+58VsyJlkEj2bcpdytmp9MmDKZkKx8OIE=",
    version = "v1.1.1",
)

go_repository(
    name = "org_modernc_memory",
    importpath = "modernc.org/memory",
    sum = "h1:utMBrFcpnQDdNsmM6asmyH/FM9TqLPS7XF7otpJmrwM=",
    version = "v1.0.4",
)

go_repository(
    name = "org_modernc_sqlite",
    importpath = "modernc.org/sqlite",
    sum = "h1:xr+yDm5hNzINGa4obXy+eq/vr3iYxpeCrHkkwTrqHT0=",
    version = "v1.0.0",
)

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()
