/*
 * helloworld_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package helloworld_pkg

import "helloworld_lib/models_pkg"
import "helloworld_lib/configuration_pkg"

/*
 * Interface for the HELLOWORLD_IMPL
 */
type HELLOWORLD interface {
    GetHelloworldGET () (*models_pkg.HelloWorldResponse, error)
}

/*
 * Factory for the HELLOWORLD interaface returning HELLOWORLD_IMPL
 */
func NewHELLOWORLD(config configuration_pkg.CONFIGURATION) *HELLOWORLD_IMPL {
    client := new(HELLOWORLD_IMPL)
    client.config = config
    return client
}