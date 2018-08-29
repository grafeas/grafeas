// Package pkg implements functions to validate that the fields of package entities being passed
// into the API meet our requirements.
package pkg

import (
	"errors"
	"fmt"

	ppb "github.com/grafeas/grafeas/proto/v1beta1/package_go_proto"
)

// ValidatePackage validates that a package has all its required fields filled in.
func ValidatePackage(p *ppb.Package) []error {
	errs := []error{}

	if p.GetName() == "" {
		errs = append(errs, errors.New("name is required"))
	}

	for i, d := range p.GetDistribution() {
		if d == nil {
			errs = append(errs, fmt.Errorf("distribution[%d] distribution cannot be null", i))
		} else {
			for _, err := range validateDistribution(d) {
				errs = append(errs, fmt.Errorf("distribution[%d].%s", i, err))
			}
		}
	}

	return errs
}

func validateDistribution(d *ppb.Distribution) []error {
	errs := []error{}

	if d.GetCpeUri() == "" {
		errs = append(errs, errors.New("cpe_uri is required"))
	}

	if ver := d.GetLatestVersion(); ver != nil {
		for _, err := range ValidateVersion(ver) {
			errs = append(errs, fmt.Errorf("version.%s", err))
		}
	}

	return errs
}

// ValidateVersion validates that a version has all its required fields filled in.
func ValidateVersion(v *ppb.Version) []error {
	errs := []error{}

	// MAXIMUM and MINIMUM version kinds are valid without a Name
	if v.GetKind() == ppb.Version_NORMAL && v.GetName() == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if v.GetKind() == ppb.Version_VERSION_KIND_UNSPECIFIED {
		errs = append(errs, errors.New("kind is required"))
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *ppb.Details) []error {
	errs := []error{}

	if i := d.GetInstallation(); i == nil {
		errs = append(errs, errors.New("installation is required"))
	} else {
		for _, err := range validateInstallation(i) {
			errs = append(errs, fmt.Errorf("installation.%s", err))
		}
	}

	return errs
}

func validateInstallation(i *ppb.Installation) []error {
	errs := []error{}

	if loc := i.GetLocation(); loc == nil {
		errs = append(errs, errors.New("location is required"))
	} else if len(loc) == 0 {
		errs = append(errs, errors.New("location requires at least 1 element"))
	} else {
		for i, l := range loc {
			if l == nil {
				errs = append(errs, fmt.Errorf("location[%d] location cannot be null", i))
			} else {
				for _, err := range validateLocation(l) {
					errs = append(errs, fmt.Errorf("location[%d].%s", i, err))
				}
			}
		}
	}

	return errs
}

func validateLocation(l *ppb.Location) []error {
	errs := []error{}

	if l.GetCpeUri() == "" {
		errs = append(errs, errors.New("cpe_uri is required"))
	}

	if v := l.GetVersion(); v != nil {
		for _, err := range ValidateVersion(v) {
			errs = append(errs, fmt.Errorf("version.%s", err))
		}
	}

	return errs
}
