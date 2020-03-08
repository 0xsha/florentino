# Florentino; Fast Static File Analysis Framework

![Florentino](https://www.arenaofvalor.com/images/heroes/skin/52101_big.jpg)

## Story
Florentino is named after a [fiction](http://arenaofvalor.com/web2017/heroDetails.html?id=521) warrior. 

Flarentino: `"I'd wear a fedora but they haven't invented them yet"`

As the sole heir to the House of Perfume, Florentino's romantic adventures were as well-known as his lavish balls [...](http://arenaofvalor.com/web2017/heroDetails.html?id=521).

Florentino: `"Ah... relationships are such a bother"`


 # Introduction 
 
Florentino is a cross-platform file analysis framework. useful for extracting static resources from malwares and unknown file analysis. 

He can help malware analysts and security researchers to quickly get a glance at an unknown file. He can't win a big war alone, though; that's why he calls for his friends to help fighting bad guys.
so he calls these friends (credits):
- [Golang](https://github.com/golang)
- [D.I.E](https://github.com/horsicq/Detect-It-Easy)
- [iocextract](https://github.com/InQuest/python-iocextract)
- [VirusTotal](https://github.com/VirusTotal)
- [Floss](https://github.com/fireeye/flare-floss) 
- [Strings](https://linux.die.net/man/1/strings) 

Without them, it was a lost war from beginning.

# Motivation
Anytime we want to analyze an unknown file, there are a couple of steps which are almost identical Florentino aims to automate some of these boring steps so an analyst can move faster with manual and dynamic analysis. 

Florentino: `"Flowers, women – I desire all that is beautiful."`

# Features 
Florentino is written in go, and it's fast!. You can run it before any other tool in your chain to gain a good grasp of your target file.
Most of the time, it's all you need to determine if a file is malicious or not!

1- File detection engine

Thanks to D.I.E, Florentino can detect hundreds of file types.

```Number of Binary signatures: 248
Number of com signatures: 200
Number of Text signatures: 14
Number of com signatures: 3
Number of MSDOS signatures: 306
Number of PE/PE+ signatures: 525
Number of DS signatures: 19
Number of EP signatures: 3
Number of ELF/ELF64 signatures: 16
Number of MACH/MACH64 signatures: 8
Total signatures: 1117 

````

Beside file detection, entropy and packer detection also performed. 


2- Scan engine 

Florentino can work various sources to analyze the file. 

- VirusTotal: we check it  for an existing report 
- Strings and IOC scan: Florentino takes it; further it will extract, scan and possibly deobfuscate strings from binary files
- Binary scan:  Florentino can work with PE x86/x64, Macho x86/x64, ELF x86/x64 files it will obtain imported symbol and libraries

3- Packer detection and unpacking 
- Currently only support PE x86 Files 
- unpack engine : [unpac.me](unpac.me)

4- Report
- All reports are stored as a text file in /data directory


Please note Florentino is not a reversing suite and its only aim is only to fasten the first analysis 
Florentino is modular and easy to extend with your own tools.

Flarentino: `Fairest ladies, my lips are like whatever I finish this later ...`

# Version

1.0.1-alpha

# Installation and Usage 

Flarentino : `"You have bad form my friend."`

check out documentation at [/docs/README.md](./docs/README.md)

# Action time: Florentino VS Ryuk Ransomware

Let's run Florentino against the trending millions dollar ransomware called [Ryuk](https://malware.wikia.org/wiki/Ryuk). 

[![asciicast](https://asciinema.org/a/OfkrF5PkylNKPl8EW36nDaf8n.svg)](https://asciinema.org/a/OfkrF5PkylNKPl8EW36nDaf8n)

`
Florentino -f 8d3f68b16f0710f858d8c1d2c699260e6f43161a5510abb0e7ba567bd7.exe
`


After one minutes or so we check /data folder


```json

{
    "detects": [
        {
            "filetype": "PE+(64)",
            "name": "Microsoft Visual C/C++(2015 v.14.0)[-]",
            "type": "compiler"
        },
        {
            "filetype": "PE+(64)",
            "name": "Microsoft Linker(14.0, Visual Studio 2015 14.0*)[EXE64]",
            "type": "linker"
        }
    ],
    "entropy": "6.07306",
    "filename": "/malwares/8d3f68b16f0710f858d8c1d2c699260e6f43161a5510abb0e7ba567bd7.exe"
}
```



```
/C REG ADD "HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /v "svchos" /t REG_SZ /d " 
```

```
Gentlemen!
Your business is at serious risk. BLAH BLAH BLAH
15RLWdVnY5n1n7mTvU1zjg67wt86dhYqNj
.....
```
- Now in less than 3 minutes, we already know its ransomware, it's not packed, we decrypted the first layer of obfuscated strings, and we already even extracted the persistence method. 
- Please consider this is NOT ready for production, the main point of releasing this is to show you how you can achieve similar results. the code can greatly improve.




# How to contribute
Florentino : `"HaHa, A wonderful day for a duel among gentlemen."`
- Add a module or fix something and then pull request.
- The endless possibility of improvements:
    - Add a web UI
    - Connect it to a Relational/NoSQL database
    - Parse each binary to its deepest details
    - Integrate r2 as provide disassembles 
    - ...
- Share it with whomever you believe can use it.
- Do the extra work and share your findings with community &hearts;
- [![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/W7W112I38)

# Learn More

[Malware fight back the tale of agent tesla](https://0xsha.io/posts/malware-fight-back-the-tale-of-agent-tesla)

[Awesome Malware Analysis](https://github.com/rshipp/awesome-malware-analysis)

[Awsome Reversing](https://github.com/tylerha97/awesome-reversing)
# License

The project is licensed under the `wtfpl` license.

