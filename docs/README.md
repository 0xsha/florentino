# Usage
Florentino is straightforward to use; all you have to do is install dependencies
and setup .EVN file


- Download [D.I.E](https://github.com/horsicq/Detect-It-Easy) 
- Download [Floss](https://github.com/fireeye/flare-floss)
- `pip3 install iocextract`



# Build and Run
- `cd cmd`
- `mkdir data`
- `touch .evn`
- example .evn
    - ``` 
          DIEC_PATH=/tools/diec
          FLOSS_PATH=/tools/floss
          VIRUSTOTAL_API=YOUR_API_KEY
      ```
- `go build main`
- `Florentino -f FILE-TO-ANALYSIS`
- now data will be available in `/data` 