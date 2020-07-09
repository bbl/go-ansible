# go-ansible

Go-ansible is a lightweight Golang package to run [Ansible](https://docs.ansible.com/) commands from Go.

## Installation

Use `go-get`:

```bash
go get -u github.com/bbl/go-ansible
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/bbl/go-ansible"
)

func main() {
    // Download dependencies from requirements 
    
    err := ansible.Galaxy().
             Requirements("/path/to/requirements.yaml").
    	     RolesPath("/path/to/roles_dir").
             Install()
    if err != nil {
        panic(err)
    }   

    // Run a playbook with extra vars
    vars := map[string]string{
        "test_key": "test_value",
    }
	pb := ansible.Playbook().
    		Inventory("/path/to/inventory.ini").
    		Path("/path/to/inventory.yaml").
    		ExtraVars(vars).
    		ExtraVars(map[string]string{"key": "value"})
    err := pb.Run()
    if err != nil {
        panic(err)
    }
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
