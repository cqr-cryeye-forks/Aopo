package common

import "flag"

func Flag(Info *Flaginfo) {
	flag.StringVar(&Info.Ip, "ip", "", "Specify IP, supported formats： 192.168.1.1｜192.168.1.1/24｜192.168.1.1/16｜192.168.1.1/8｜192.168.1.1,192.168.1.2")
	flag.StringVar(&Info.Ipfile, "ipf", "", "Read the target from the file, the supported format is as above. Note: one per line")
	flag.BoolVar(&Info.All, "all", false, "Automatic detection of intranet assets, default scan: 10A|172A|192B")
	flag.StringVar(&Info.Passdic, "addpass", "", "Custom dictionaries. Merging into the built-in dictionary library")
	flag.StringVar(&Info.JsonFile, "json", "results.json", "Save results to results.json")
	//flag.IntVar(&Info.Threads, "t", 500, "自定义线程 默认500线程")
	flag.Parse()
}
