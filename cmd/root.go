package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var templateFile, propertiesFile string

var rootCmd = &cobra.Command{
	Use:   "ninja_jinja",
	Short: "Tell me why",
	Run: func(cmd *cobra.Command, args []string) {
		validateInput()

		viper.SetConfigFile(propertiesFile) // name of config file (without extension)
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}

		props := map[string]string{}

		for _, k := range viper.AllKeys() {
			props[k] = viper.GetString(k)
			// fmt.Println("Key:", k, "Value:", props[k])

		}
		tmpl := template.Must(template.ParseFiles(templateFile))
		err = tmpl.Execute(os.Stdout, props)
		if err != nil {
			panic(err)
		}
	},
}

// Execute tanginang warning yan
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&templateFile, "template", "", "ain't nothing but a template.")
	rootCmd.PersistentFlags().StringVar(&propertiesFile, "properties", "", "ain't nothing but a mistake.")
	rootCmd.MarkPersistentFlagRequired("template")
	rootCmd.MarkPersistentFlagRequired("properties")

}

func validateInput() {
	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		fmt.Printf("Template: %s does not exist", templateFile)
		os.Exit(1)
	}
	if _, err := os.Stat(propertiesFile); os.IsNotExist(err) {
		fmt.Printf("Properties: %s does not exist", propertiesFile)
		os.Exit(1)
	}
}
