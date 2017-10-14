// Copyright 2017 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import "github.com/grafeas/grafeas/server-go/api/v1alpha1"

func Occurrence(noteName string) api.Occurrence {
	return api.Occurrence{
		Name:        "projects/test-project/occurrences/134",
		ResourceUrl: "gcr.io/foo/bar",
		NoteName:    noteName,
		Kind:        "PACKAGE_VULNERABILITY",
		VulnerabilityDetails: api.VulnerabilityDetails{
			Severity:  "HIGH",
			CvssScore: 7.5,
			PackageIssue: []api.PackageIssue{
				{
					SeverityName: "HIGH",
					AffectedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: api.Version{
							Name:     "52.1",
							Revision: "8+deb8u3",
						},
					},
					FixedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: api.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
			},
		},
	}
}

func Note() api.Note {
	return api.Note{
		Name:             "projects/vulnerability-scanner-a/notes/CVE-1999-0710",
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             "PACKAGE_VULNERABILITY",
		VulnerabilityType: api.VulnerabilityType{
			CvssScore: 7.5,
			Severity:  "HIGH",
			Details: []api.Detail{
				{
					CpeUri:   "cpe:/o:debian:debian_linux:7",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: api.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:7",
						Package_: "icu",
						Version: api.Version{
							Name:     "4.8.1.1",
							Revision: "12+deb7u6",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:8",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: api.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: api.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:9",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: api.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:9",
						Package_: "icu",
						Version: api.Version{
							Name:     "55.1",
							Revision: "3",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
					Package_: "andriod",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: api.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "MEDIUM",

					FixedLocation: api.VulnerabilityLocation{
						CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
						Package_: "andriod",
						Version: api.Version{
							Kind: "MAXIMUM",
						},
					},
				},
			},
		},
		RelatedUrl: []api.RelatedUrl{
			{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}

func Operation() api.Operation {
	return api.Operation{
		Name:     "projects/vulnerability-scanner-a/operations/foo",
		Metadata: map[string]string{"StartTime": "0916162344"},
		Done:     false,
	}
}
