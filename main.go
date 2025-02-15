package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// https://github.com/pdfcpu/pdfcpu/blob/4a2f04f3a268f556a4e3145d3978690b7b559cc7/pkg/api/test/encryption_test.go#L102
func main() {
	conf = confForAlgorithm(aes, keyLength, "upw", "opw")
	if err = api.ChangeUserPasswordFile(outFile, "", "upw", "upwNew", conf); err != nil {
		t.Fatalf("%s: change upw %s: %v\n", msg, outFile, err)
	}

	// Change owner password.
	conf = confForAlgorithm(aes, keyLength, "upwNew", "opw")
	if err = api.ChangeOwnerPasswordFile(outFile, "", "opw", "opwNew", conf); err != nil {
		t.Fatalf("%s: change opw %s: %v\n", msg, outFile, err)
	}
}

// func processListViewerPreferencesCommand(conf *model.Configuration) {
// 	if len(flag.Args()) != 1 || selectedPages != "" {
// 		fmt.Fprintf(os.Stderr, "usage: %s\n", usageViewerPreferencesList)
// 		os.Exit(1)
// 	}

// 	inFile := flag.Arg(0)
// 	// if conf.CheckFileNameExt {
// 	// 	ensurePDFExtension(inFile)
// 	// }

// 	if json {
// 		log.SetCLILogger(nil)
// 	}

// 	process(cli.ListViewerPreferencesCommand(inFile, all, json, conf))
// }

// // ListViewerPreferencesCommand creates a new command to list the viewer preferences.
// func ListViewerPreferencesCommand(inFile string, all, json bool, conf *model.Configuration) *Command {

// 	if conf == nil {
// 		conf = model.NewDefaultConfiguration()
// 	}
// 	conf.Cmd = model.LISTVIEWERPREFERENCES
// 	return &Command{
// 		Mode:     model.LISTVIEWERPREFERENCES,
// 		InFile:   &inFile,
// 		BoolVal1: all,
// 		BoolVal2: json,
// 		Conf:     conf}
// }

// func process(cmd *cli.Command) {
// 	out, err := cli.Process(cmd)
// 	if err != nil {
// 		if needStackTrace {
// 			fmt.Fprintf(os.Stderr, "Fatal: %+v\n", err)
// 		} else {
// 			fmt.Fprintf(os.Stderr, "%v\n", err)
// 		}
// 		os.Exit(1)
// 	}

// 	if out != nil && !quiet {
// 		for _, s := range out {
// 			fmt.Fprintln(os.Stdout, s)
// 		}
// 	}
// 	//os.Exit(0)
// }

// // Process executes a pdfcpu command.
// func Process(cmd *Command) (out []string, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			err = errors.Errorf("unexpected panic attack: %v\n", r)
// 		}
// 	}()

// 	cmd.Conf.Cmd = cmd.Mode

// 	if f, ok := cmdMap[cmd.Mode]; ok {
// 		return f(cmd)
// 	}

// 	return nil, errors.Errorf("pdfcpu: process: Unknown command mode %d\n", cmd.Mode)
// }
