// Package deployment implements functions to validate that the fields of deployment entities being
// passed into the API meet our requirements.
package deployment

import (
	"errors"
	"fmt"

	dpb "github.com/grafeas/grafeas/proto/v1beta1/deployment_go_proto"
)

// ValidateDeployable validates that a deployable has all its required fields filled in.
func ValidateDeployable(d *dpb.Deployable) []error {
	errs := []error{}

	if r := d.GetResourceUri(); r == nil {
		errs = append(errs, errors.New("resource_uri is required"))
	} else if len(r) == 0 {
		errs = append(errs, errors.New("resource_uri requires at least 1 element"))
	} else {
		for i, r := range r {
			if r == "" {
				errs = append(errs, fmt.Errorf("resource_uri[%d] cannot be empty", i))
			}
		}
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *dpb.Details) []error {
	errs := []error{}

	if dp := d.GetDeployment(); dp == nil {
		errs = append(errs, errors.New("deployment is required"))
	} else {
		for _, err := range validateDeployment(dp) {
			errs = append(errs, fmt.Errorf("deployment.%s", err))
		}
	}

	return errs
}

func validateDeployment(d *dpb.Deployment) []error {
	errs := []error{}

	if d.GetDeployTime() == nil {
		errs = append(errs, errors.New("deploy_time is required"))
	}

	return errs
}
