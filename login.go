// +build darwin
// +build linux 

package main

import (
	"fmt"
	// "log"
	"os/exec"
)

// if vault not found in path then downlaod vault binary and extract it 

// const VAULT_ADDR="https://lookup.vimaan.app"
const SSH_KEY_LOCATION="$HOME/.ssh/id_rsa"
const VAULT_DOWNLOAD_LINK="https://releases.hashicorp.com/vault/1.15.5/vault_1.15.5_linux_arm64.zip"

// download package using go
// unzip using go
func main(){
    vaultLogin()
}

func vaultLogin() {
    command_to_check_if_vault_exists := "vault"
    cmd_check := exec.Command("sh", "-c", command_to_check_if_vault_exists)
    output_check, err_check := cmd_check.CombinedOutput()
    if err_check != nil {
        // downloadVault()
        fmt.Println("vault not found in path. Downloading vault")
    }
    fmt.Printf("vault found in path. Output:\n%s\n", string(output_check))

    // command_to_download := "wget "+VAULT_DOWNLOAD_LINK+" -O vault.zip"
	// cmd_download := exec.Command( "sh", "-c", command_to_download )

    // command_to_unzip := "unzip vault.zip"
    // cmd_unzip := exec.Command("sh", "-c", command_to_unzip)
	
    // output_download, err_download := cmd_download.CombinedOutput()
    // output_unzip, err_unzip := cmd_unzip.CombinedOutput()

	// if err_download != nil {
	// 	log.Fatalf("vault login failed with %s\n", err_download)
	// }
    // if err_unzip != nil {
    //     log.Fatalf("vault login failed with %s\n", err_unzip)
    // }

    // fmt.Printf("vault download successful. Output:\n%s\n", string(output_download))
    // fmt.Printf("unzip successfull.", string(output_unzip))
}
