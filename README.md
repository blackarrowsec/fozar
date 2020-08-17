# Fozar

[![](https://img.shields.io/badge/Category-Data%20gathering-E5A505?style=flat-square)]() [![](https://img.shields.io/badge/Language-Golang-E5A505?style=flat-square)]()


**Fozar** allows you to traverse commits across multiple repositories matching against user supplied regex. It also allows for the user to set exclusions on certain files for them not to be analyzed.

As the tool is written in **Golang**, it can be used in both Windows and Linux.

## Installation


```shell
git clone https://github.com/blackarrowsec/fozar.git
cd fozar/
go build
```

## Usage

Launch the tool against a directory tree with one or more git repositories on it.

```
PS C:\Pentest\Tools\fozar> .\fozar.exe -h
 _____
|  ___|___  ____ __ _  _ __
| |_  / _ \|_  // _` || '__|
|  _|| (_) |/ /| (_| || |
|_|   \___//___|\__,_||_|
                        By @30vh1 [https://blackarrow.net] [https://tarlogic.com]
  -A int
        Following lines to show on match
  -B int
        Previous lines to show on match
  -config string
        Yaml config location (default "config/config.yml")
  -debug
        Print debug information (slower!)
  -of string
        Filename for the html file output
  -ot string
        Filename for the Markdown file output
  -path string
        Folder from which to start searching
```

There are a couple of scripts under the `script/` folder. You can pipe URLs to them in order to download batch repositories.
The scripts will create an `output/` folder with the raw repositories under `output/raw/` and the actual repositories under `output/repo/` (this is the folder which you wanna use Fozar on)

### Configuration

The configuration file `config/config.yml` keeps two lists of regular expressions. The first one known as **rules** keeps all the matches you want check against whilst the rules under the **avoid** section are matched against file names to avoid analyzing them.

```yaml
rules:
    - '[pP][aA][sS][sS][wW][oO][rR][dD]\s*=\s*"' # This is an inline comment
    - '"access_token":'
    - '[pP][aA][sS][sS]\s+=\s+?"'
    - ...
avoid: 
    - '.exe'
    - '.war'
    - '.rar'
    - ...
```
### Output Modes

#### HTML report

The HTML report generates an **easy on the eye** output on HTML format. When selecting this output it is necessary to have the `templates/` folder on the same directory from which **Fozar** is being run.

![image-20200813121135660](https://raw.githubusercontent.com/blackarrowsec/fozar/master/resources/fozar_example_fancy_report.gif)

#### Text Report

The text report generates a Markdown file. This output doesn't have any special requirements. A tool such is [Typora](https://typora.io) is highly recommended for the Markdown output visualization.

* Plain Markdown text file

![image-20200813121044260](https://raw.githubusercontent.com/blackarrowsec/fozar/master/resources/fozar_example_text_report_plain.png)

* Interpreted Markdown text file

![image-20200813120937214](https://raw.githubusercontent.com/blackarrowsec/fozar/master/resources/fozar_example_text_report_interpreted.png)
## Examples

For getting in touch with the tool you can try the following commands. The output result will be on `output/Fozar Report/`

* **Windows**

```powershell
cat .\scripts\test.txt | .\scripts\bulk_clone.ps1
.\fozar.exe -A 3 -B 3 -config .\config\config.yml -of test.html -path .\output\repo\
```

* **Linux** 

```bash
cat ./scripts/test.txt | ./scripts/bulk_clone.sh
./fozar -A 3 -B 3 -config ./config/config.yml -of test.html -path ./output/repo/
```

## Author

Marcos Carro ([@30vh1](https://github.com/30vh1)) [ [www.blackarrow.net](http://blackarrow.net/) - [www.tarlogic.com](https://www.tarlogic.com/en/) ]


## License
All the code included in this project is licensed under the terms of the GNU AGPLv3 license.

#

[![](https://img.shields.io/badge/www-blackarrow.net-E5A505?style=flat-square)](https://www.blackarrow.net) [![](https://img.shields.io/badge/twitter-@BlackArrowSec-00aced?style=flat-square&logo=twitter&logoColor=white)](https://twitter.com/BlackArrowSec) [![](https://img.shields.io/badge/linkedin-@BlackArrowSec-0084b4?style=flat-square&logo=linkedin&logoColor=white)](https://www.linkedin.com/company/blackarrowsec/)
