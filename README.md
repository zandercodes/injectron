<p align="center">
  <img src="assets/logo.svg" alt="Injectron Logo" width="200" height="200" />
</p>

<p align="center">
Injectron another way to inject Electron Apps
</br>
</br>
</p>

<div align="center">
<strong>
<samp>

[![GitHub](https://img.shields.io/github/license/zandercodes/injectron)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zandercodes/injectron)](https://goreportcard.com/report/github.com/zandercodes/injectron)
[![Go Docs](https://pkg.go.dev/badge/github.com/zandercodes/injectron)](https://pkg.go.dev/github.com/zandercodes/injectron)
[![GitHub contributors](https://img.shields.io/github/contributors/zandercodes/injectron?color=green)](https://github.com/zandercodes/injectron/graphs/contributors)
[![GitHub last commit](https://img.shields.io/github/last-commit/zandercodes/injectron)](https://github.com/zandercodes/injectron/commits)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/zandercodes/injectron)](https://go.dev/)
[![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/zandercodes/injectron?label=version)](https://github.com/zandercodes/injectron/releases)

</samp>
</strong>
</div>

## Table of Contents
- [Work in progress](#work-in-progress)
- [What is Injectron?](#what-is-injectron)
- [How to use?](#how-to-use)
- [How to install?](#how-to-install)
- [How to build?](#how-to-build)
- [How to contribute?](#how-to-contribute)
- [How to donate?](#how-to-donate)
- [License](#license)

## Work in progress
The project is still in development.  
It is not recommended to use it in production.  
Not all features are implemented yet.  
Not working at the moment.

## What is Injectron?
Injectron is an JavaScript injector for Electron Apps. It allows you to inject JavaScript code into the Electron App. It is useful for creating and Modify at Runtime.

## How to use?
Injectron is very easy to use. You need one or more JavaScript files that you want to insert into the application. The inject function takes two arguments, the first is the executable file and the second is the JavaScript file to be loaded into the program.

```
$ injectron inject <executable> <javascript>
```

## How to install?
You can download the latest version of Injectron from the [releases](https://github.com/zandercodes/injectron/releases) page.

## How to build?
You can build Injectron from source. You need to install [Go](https://go.dev/) to build Injectron. You can build Injectron with the following commands.

```
$ go build .
```

## How to contribute?
You can contribute to Injectron by creating a pull request. You can also contribute by reporting bugs and suggesting new features. You can also contribute by donating to the project.

## How to donate?
You can donate at Ko-fi or you can also donate via Bitcoin.  
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/G2G6M6UL4)  
Bitcoin: `bc1qar5tlrhq7zmj5xdyc73jtzwpj20vx4x2yhhdfc`

## License
Injectron is licensed under the [GNU GENERAL PUBLIC LICENSE](LICENSE).