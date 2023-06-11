# Injectron another way to inject Electron Apps

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
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/)  
Bitcoin: `bc1qar5tlrhq7zmj5xdyc73jtzwpj20vx4x2yhhdfc`

## License
Injectron is licensed under the [GNU GENERAL PUBLIC LICENSE](LICENSE).