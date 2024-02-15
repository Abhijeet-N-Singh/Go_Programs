Go for Visual Studio Code
Slack

The VS Code Go extension provides rich language support for the Go programming language.

Requirements
Visual Studio Code 1.75 or newer (or editors compatible with VS Code 1.75+ APIs)
Go 1.18 or newer
Quick Start
Welcome! 👋🏻
Whether you are new to Go or an experienced Go developer, we hope this extension fits your needs and enhances your development experience.

Install Go 1.18 or newer if you haven't already.

Install the VS Code Go extension.

Open any Go file or go.mod file to automatically activate the extension. The Go status bar appears in the bottom right corner of the window and displays your Go version.

The extension depends on go, gopls (the Go language server), and optional tools depending on your settings. If gopls is missing, the extension will try to install it. The :zap: sign next to the Go version indicates the language server is running, and you are ready to go.


(Install Missing Tools)

You are ready to Go :-)    🎉🎉🎉

What's next
Explore more features of the VS Code Go extension.
View the settings documentation and advanced topics to customize the extension.
View the tools documentation for a complete list of tools the VS Code Go extension depends on. You can install additional tools and update them by using "Go: Install/Update Tools".
Solve issues with the general troubleshooting and debugging troubleshooting guides.
file an issue for problems with the extension.
Start a GitHub discussion or get help on Stack Overflow.
Explore Go language resources on go.dev/learn and golang.org/help.
If you are new to Go, this article provides the overview on Go code organization and basic go commands. Watch "Getting started with VS Code Go" for an explanation of how to build your first Go application using VS Code Go.

Feature highlights
IntelliSense - Results appear for symbols as you type.
Code navigation - Jump to or peek at a symbol's declaration.
Code editing - Support for saved snippets, formatting and code organization, and automatic organization of imports.
Diagnostics - Build, vet, and lint errors shown as you type or on save.
Enhanced support for testing and debugging
See the full feature breakdown for more details.


(Code completion and Signature Help)

In addition to integrated editing features, the extension provides several commands for working with Go files. You can access any of these by opening the Command Palette (Ctrl+Shift+P on Linux/Windows and Cmd+Shift+P on Mac), and then typing in the command name. See the full list of commands provided by this extension.


(Toggle Test File)

For better syntax highlighting, we recommend enabling semantic highlighting by turning on Gopls' ui.semanticTokens setting. "gopls": { "ui.semanticTokens": true }

Setting up your workspace
The VS Code Go extension supports both GOPATH and Go modules modes.

Go modules are used to manage dependencies in recent versions of Go. Modules replace the GOPATH-based approach to specifying which source files are used in a given build, and they are the default build mode in go1.16+. We highly recommend Go development in module mode. If you are working on existing projects, please consider migrating to modules.

Unlike the traditional GOPATH mode, module mode does not require the workspace to be located under GOPATH nor to use a specific structure. A module is defined by a directory tree of Go source files with a go.mod file in the tree's root directory.

Your project may involve one or more modules. If you are working with multiple modules or uncommon project layouts, you will need to configure your workspace by using Workspace Folders. See the Supported workspace layouts documentation for more information.

Preview version
If you'd like to get early access to new features and bug fixes, you can use the nightly build of this extension. Learn how to install it in by reading the Go Nightly documentation.

Telemetry
VS Code Go extension relies on the Go Telemetry to learn insights about the performance and stability of the extension and the language server (gopls). Go Telemetry data uploading is disabled by default and can be enabled with the following command:

go run golang.org/x/telemetry/cmd/gotelemetry@latest on
After telemetry is enabled, the language server will upload metrics and stack traces to telemetry.go.dev. You can inspect what data is collected and can be uploaded by running:

go run golang.org/x/telemetry/cmd/gotelemetry@latest view
If we get enough adoption, this data can significantly advance the pace of the Go extension development, and help us meet a higher standard of reliability. For example:

Even with semi-automated crash reports in VS Code, we've seen several crashers go unreported for weeks or months.
Even with a suite of benchmarks, some performance regressions don't show up in our benchmark environment (such as the completion bug mentioned below!).
Even with lots of great ideas for how to improve gopls, we have limited resources. Telemetry can help us identify which new features are most important, and which existing features aren't being used or aren't working well.
These are just a few ways that telemetry can improve gopls. The telemetry blog post series contains many more.

Go telemetry is designed to be transparent and privacy-preserving. If you have concerns about enabling telemetry, you can learn more at https://telemetry.go.dev/privacy.

Contributing
We welcome your contributions and thank you for working to improve the Go development experience in VS Code. If you would like to help work on the VS Code Go extension, see our contribution guide to learn how to build and run the VS Code Go extension locally and contribute to the project.

Code of Conduct
This project follows the Go Community Code of Conduct. If you encounter a conduct-related issue, please mail conduct@golang.org.

License
MIT

Categories
Programming LanguagesSnippetsLintersDebuggersFormattersTesting
Tags
debuggersgoGo Checksum FileGo Module FileGo Template FileGo Work Filego.modgo.sumgo.workgolanggoplsgotmplmulti-root readysnippet
Works with
Universal
