// +build darwin

package main

import (
	"fmt"
	"log"
	"os"
    "os/exec"
)

const VAULT_ADDR="https://lookup.vimaan.app"
const SSH_KEY_LOCATION="$HOME/.ssh/id_rsa"
const VAULT_DOWNLOAD_LINK="https://releases.hashicorp.com/vault/1.15.5/vault_1.15.5_"
const SSH_SIGNING_ENGINE="ssh-test-blrvm2"
// download package using go
// unzip using go
func main(){
    vaultNeedsToBeDownloaded := check_if_vault_exists()
    if vaultNeedsToBeDownloaded == -1{
            downloadVault()
            unzipVault()
        }
        vaultLogin()
        signSshKey()
    }

    func checkDist()(string){
        command_to_check_dist := "uname -s | tr 'A-Z' 'a-z'"
        cmd_check_dist := exec.Command("sh", "-c", command_to_check_dist)
        output_check_dist, err_check_dist := cmd_check_dist.CombinedOutput()
        if err_check_dist != nil {
            log.Fatalf("checkDist failed with %s\n", err_check_dist)
        }
        fmt.Printf("checkDist successful. Output:\n%s\n", string(output_check_dist))
        return string(output_check_dist)
    }

    func checkArch()(string){
        command_to_check_arch := "uname -m | tr 'A-Z' 'a-z'"
        cmd_check_arch := exec.Command("sh", "-c", command_to_check_arch)
        output_check_arch, err_check_arch := cmd_check_arch.CombinedOutput()
        if err_check_arch != nil {
            log.Fatalf("checkArch failed with %s\n", err_check_arch)
        }
        fmt.Printf("checkArch successful. Output:\n%s\n", string(output_check_arch))
        return string(output_check_arch)
    }
    
    func unzipVault(){
        command_to_unzip := "unzip vault.zip"
        cmd_unzip := exec.Command("sh", "-c", command_to_unzip)
        output_unzip, err_unzip := cmd_unzip.CombinedOutput()
        if err_unzip != nil {
            log.Fatalf("vault unzipping failed with %s\n", err_unzip)
        }
        fmt.Printf("unzip successfull.", string(output_unzip))
    }
    
    func downloadVault(){
        arch := checkArch()
        dist := checkDist()
        downloadLink := VAULT_DOWNLOAD_LINK+dist+"_"+arch+".zip"
        fmt.Printf(downloadLink)
        command_to_download := "curl "+downloadLink+" -o vault.zip"
        cmd_download := exec.Command( "sh", "-c", command_to_download )
        output_download, err_download := cmd_download.CombinedOutput()
        if err_download != nil {
            log.Fatalf("vault login failed with %s\n", err_download)
        }
        fmt.Printf("vault download successful. Output:\n%s\n", string(output_download))
    }
    
    func check_if_vault_exists()(int){
        path := os.Getenv("PATH")
        cwd := os.Getenv("PWD")
        newPath := path + ":" + cwd
    os.Setenv("PATH", newPath)
    // fmt.Printf("new path: %s", os.Getenv("PATH"))

    command_to_check_if_vault_exists := "which vault"
    cmd_check := exec.Command("sh", "-c", command_to_check_if_vault_exists)
    output_check, err_check := cmd_check.CombinedOutput()
    if err_check != nil {
        fmt.Println("vault not found in path. Will Download vault")
        return -1
    }
    fmt.Printf("vault found in path. Output:\n%s\n", string(output_check))
    return 1
}

func vaultLogin() {
    loginCommandString := "vault login -method=oidc 1>/dev/null"
    loginCommand := exec.Command("sh", "-c", loginCommandString)
    // fmt.Printf(os.Getenv("PATH"))
    os.Setenv("VAULT_ADDR", VAULT_ADDR)
    outputLogin, errLogin := loginCommand.CombinedOutput()
    if errLogin != nil {
        log.Fatalf("vault login failed with %s\n", errLogin)
    }
    fmt.Printf("vault login successful. Output:\n%s\n", string(outputLogin))
}

func signSshKey(){
    signCommandString := "vault write -field=signed_key ssh-test-blrvm2/sign/blrvm-developer public_key=@$HOME/.ssh/id_rsa.pub > id_rsa-cert.pub"
    signCommand := exec.Command("sh", "-c", signCommandString)
    outputSign, errSign := signCommand.CombinedOutput()
    if errSign != nil {
        log.Fatalf("ssh key signing failed with %s\n", errSign)
    }
    fmt.Printf("ssh key signing successful. Output:\n%s\n", string(outputSign))
}