<div id="top"></div>

[![Go Reference](https://pkg.go.dev/badge/github.com/axpira/goplogadapter.svg)](https://pkg.go.dev/github.com/axpira/goplogadapter)
[![Go Report Card](https://goreportcard.com/badge/github.com/axpira/goplogadapter)](https://goreportcard.com/report/github.com/axpira/goplogadapter)
[![Coverage](http://gocover.io/_badge/github.com/axpira/goplogadapter)](http://gocover.io/github.com/axpira/goplogadapter)

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">GOP Log Adapter</h3>

  <p align="center">
    Adapter to use with <a href="https://github.com/axpira/gop">Gop Log</a>

    You must just implement one method to create your own log implementation

    Or choose one of implementations like logrus or json
    <br />
    <br />
    <a href="https://github.com/axpira/goplogadapter/issues">Report Bug</a>

    <a href="https://github.com/axpira/goplogadapter/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


This project implements [gop log](https://github.com/axpira/gop) and expose just one function:
```go
type FormatterFunc func(log.Level, string, error, map[string]interface{})
```

This function will be called every time need to send something to log, and you need just to implements and set as default _Formatter_

In this repository has some sub folder, for each sub folder is one of implementation you can use, like:
* json: Just convert to json and send to output
* logrus: Use [logrus](https://github.com/sirupsen/logrus) to send log

All this are sub module, so if you want to use you will need to import in your project

<p align="right">(<a href="#top">back to top</a>)</p>


### Built With

* [Go lang](https://golang.org/)


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Installation

```bash
go get -u github.com/axpira/goplogadapter
```
Or choosing one implementation:
```bash
go get -u github.com/axpira/goplogadapter/logrus
```

## Getting Started

You will need to use [gop log](https://github.com/axpira/gop) to log anything.

And to choose this implementation just do it:
```go
import _ "github.com/axpira/goplogadapter/logrus"
```

Now when you log:
```go
import (
	"github.com/axpira/gop/log"
	_ "github.com/axpira/goplogadapter/logrus"
)

func main() {
	log.Info("Hello World")

	log.Inf(log.
		Str("str_field", "hello").
		Int("int_field", 42).
		Msg("Hello World"),
	)
```

The logrus will be called to send log



<!-- USAGE EXAMPLES -->
## Usage

To create a new logger and set as default for gop:
```go
import (
	"github.com/axpira/gop/log"
	gla "github.com/axpira/goplogadapter"
)

func init() {
	MyFormatter := func(level log.Level, msg string, err error, m map[string]interface{}) {
		fmt.Printf("[%v] [%v] %s %v", level, err, msg, m)
	}
	log.DefaultLogger = gla.New(gla.WithFormatter(MyFormatter), gpa.WithLevel(log.LevelTrace))
}
```
This example is creating a new formatter,
setting as default for gop and change minimum level to log to _Trace_.

But you can just use one of implementation bellow

### Logrus

To configure anything you want please check at [logrus](https://github.com/sirupsen/logrus).
For example:
```go
import (
	"os"
	"github.com/axpira/gop/log"
	_ "github.com/axpira/goplogadapter/logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Now the log format will be JSON with caller and output to os.Stdout
	log.Inf(log.
		Str("str_field", "hello").
		Int("int_field", 42).
		Msg("Hello World"),
)
```

### Json

Using default golang json parse to send fields to output
To configure anything you want please check at [logrus](https://github.com/sirupsen/logrus).
For example:
```go
import (
	"os"
	"github.com/axpira/gop/log"
	_ "github.com/axpira/goplogadapter/json"
)

func main() {
	// Changing default output
	json.Output = os.Stderr

	log.Inf(log.
		Str("str_field", "hello").
		Int("int_field", 42).
		Msg("Hello World"),
)
```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Thiago Ferreira - thiagogbferreira@gmail.com

Project Link: [https://github.com/axpira/goplogadapter](https://github.com/axpira/goplogadapter)

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/axpira/goplogadapter.svg?style=for-the-badge
[contributors-url]: https://github.com/axpira/goplogadapter/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/axpira/goplogadapter.svg?style=for-the-badge
[forks-url]: https://github.com/axpira/goplogadapter/network/members
[stars-shield]: https://img.shields.io/github/stars/axpira/goplogadapter.svg?style=for-the-badge
[stars-url]: https://github.com/axpira/goplogadapter/stargazers
[issues-shield]: https://img.shields.io/github/issues/axpira/goplogadapter.svg?style=for-the-badge
[issues-url]: https://github.com/axpira/goplogadapter/issues
[license-shield]: https://img.shields.io/github/license/axpira/goplogadapter.svg?style=for-the-badge
[license-url]: https://github.com/axpira/goplogadapter/blob/master/LICENSE.txt
