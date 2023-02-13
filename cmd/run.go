package cmd

import (
        "fmt"
	"log"

        "github.com/spf13/cobra"

//	cosmos "github.com/node-a-team/Cosmos-IE/chains/cosmos"
//        terra "github.com/node-a-team/Cosmos-IE/chains/terra"
//        iris "github.com/node-a-team/Cosmos-IE/chains/iris"
//        kava "github.com/node-a-team/Cosmos-IE/chains/kava"
//        iov "github.com/node-a-team/Cosmos-IE/chains/iov"
//        emoney "github.com/node-a-team/Cosmos-IE/chains/emoney"
//        band "github.com/node-a-team/Cosmos-IE/chains/bandprotocol"
//        paloma "github.com/node-a-team/Cosmos-IE/chains/paloma"

	exporter "github.com/andrii1890/Cosmos-IE/exporter"

)

var (

	chainList []string

	// command로 안 받을 경우 defult 값 지정
	chain string = ""
        restAddr string = "localhost:10317"
	listenPort string = "10661"

	operAddr string= ""
)

// versionCmd represents the version command
var runCmd = &cobra.Command{
        Use:   "run",
        Short: "Validator Operator Address",
        Long: `Be sure to enter either Validator Operator Address
ex#1_Local REST Server) Cosmos-IE run --chain paloma --oper-addr palomavaloper1swa5kcf9cl5dx2ypx0c5r9e5qdfnzp9w0yj7uv
ex#2_Remote REST Server) Cosmos-IE run --chain paloma --oper-addr palomavaloper1swa5kcf9cl5dx2ypx0c5r9e5qdfnzp9w0yj7uv --rest-server 127.0.0.1:10317 
`,
        Run: func(cmd *cobra.Command, args []string) {
		check_chain()
		set_config()
		run()
        },
}

func init() {
        rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&chain, "chain", "c", "", "Chain name of the monitoring node")
        runCmd.MarkFlagRequired("chain")

	runCmd.Flags().StringVarP(&restAddr, "rest-server", "", "localhost:10317", "<host>:<port> to Rest-Server(LCD-Server) interface for the selected chain")
	runCmd.Flags().StringVarP(&listenPort, "port", "p", "10661", "Port to listen for Prometheus collector connections")
	runCmd.Flags().StringVarP(&operAddr, "oper-addr", "", "", "Operator address for Validator")
}

func check_chain() {

	// chain check
	chainCheck := false
	for _, c := range chainList {
		if chain == c {
			chainCheck = true
		}
	}
	if !chainCheck {
		log.Fatal(fmt.Sprintf("[Error] %s is not supported", chain) +fmt.Sprint("\nList of supported chains: ", chainList))
	}


	if (operAddr == "") {
		log.Fatal("[Error] --oper-addr was not entered.")
	}
}

func run() {
	switch chain {
//                case "cosmos":
//                        cosmos.Main(listenPort)
//                case "terra","emoney":
//                        exporter.Go(chain, listenPort)
//                case "iris":
//                        iris.Main(listenPort)
//                case "kava":
//                        kava.Main(listenPort)
//                case "iov":
//                        iov.Main(listenPort)
//		case "band":
//                        band.Main(listenPort)
		default:
			exporter.Go(chain, listenPort)

        }
}
