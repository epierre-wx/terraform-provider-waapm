package main


func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: waapm.Provider,
	})
}
