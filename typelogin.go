package main

type loginResponse struct {
	Status int `json:"status"`
	Data   struct {
		User struct {
			ID                    string `json:"id"`
			FirstName             string `json:"firstName"`
			LastName              string `json:"lastName"`
			Email                 string `json:"email"`
			Country               string `json:"country"`
			UILanguage            string `json:"uiLanguage"`
			CommunicationLanguage string `json:"communicationLanguage"`
			AccountType           string `json:"accountType"`
			Uom                   string `json:"uom"`
			DateFormat            string `json:"dateFormat"`
			TimeFormat            string `json:"timeFormat"`
			EmailDay              []int  `json:"emailDay"`
			System                struct {
				Messages struct {
					AppReviewBanner                  int    `json:"appReviewBanner"`
					FirstUsePhoenix                  int    `json:"firstUsePhoenix"`
					FirstUsePhoenixReportsDataMerged int    `json:"firstUsePhoenixReportsDataMerged"`
					LluGettingStartedBanner          int    `json:"lluGettingStartedBanner"`
					LluNewFeatureModal               int    `json:"lluNewFeatureModal"`
					LvWebPostRelease                 string `json:"lvWebPostRelease"`
				} `json:"messages"`
			} `json:"system"`
			Details struct {
			} `json:"details"`
			TwoFactor struct {
				PrimaryMethod   string `json:"primaryMethod"`
				PrimaryValue    string `json:"primaryValue"`
				SecondaryMethod string `json:"secondaryMethod"`
				SecondaryValue  string `json:"secondaryValue"`
			} `json:"twoFactor"`
			Created   int `json:"created"`
			LastLogin int `json:"lastLogin"`
			Programs  struct {
			} `json:"programs"`
			DateOfBirth int `json:"dateOfBirth"`
			Practices   struct {
				NAMING_FAILED struct {
					ID          string `json:"id"`
					PracticeID  string `json:"practiceId"`
					Name        string `json:"name"`
					Address1    string `json:"address1"`
					City        string `json:"city"`
					State       string `json:"state"`
					Zip         string `json:"zip"`
					PhoneNumber string `json:"phoneNumber"`
					Records     any    `json:"records"`
				} `json:""`
			} `json:"practices"`
			Devices struct {
				NAMING_FAILED struct {
					ID         string `json:"id"`
					Nickname   string `json:"nickname"`
					Sn         string `json:"sn"`
					Type       int    `json:"type"`
					UploadDate int    `json:"uploadDate"`
				} `json:""`
				NAMING_FAILED0 struct {
					ID         string `json:"id"`
					Nickname   string `json:"nickname"`
					Sn         string `json:"sn"`
					Type       int    `json:"type"`
					UploadDate int    `json:"uploadDate"`
				} `json:""`
				NAMING_FAILED1 struct {
					ID         string `json:"id"`
					Nickname   string `json:"nickname"`
					Sn         string `json:"sn"`
					Type       int    `json:"type"`
					UploadDate int    `json:"uploadDate"`
				} `json:""`
			} `json:"devices"`
			Consents struct {
				AssistLibre struct {
					PolicyAccept int `json:"policyAccept"`
					History      []struct {
						PolicyAccept int  `json:"policyAccept"`
						Declined     bool `json:"declined,omitempty"`
					} `json:"history"`
				} `json:"assistLibre"`
				Hipaa struct {
					PolicyAccept int `json:"policyAccept"`
					History      []struct {
						PolicyAccept int `json:"policyAccept"`
					} `json:"history"`
				} `json:"hipaa"`
				HipaaLibre struct {
					PolicyAccept int `json:"policyAccept"`
					History      []struct {
						PolicyAccept int `json:"policyAccept"`
					} `json:"history"`
				} `json:"hipaaLibre"`
				Llu struct {
					PolicyAccept int `json:"policyAccept"`
					TouAccept    int `json:"touAccept"`
				} `json:"llu"`
				RealWorldEvidence struct {
					PolicyAccept int `json:"policyAccept"`
					History      []struct {
						PolicyAccept int  `json:"policyAccept"`
						Declined     bool `json:"declined,omitempty"`
					} `json:"history"`
				} `json:"realWorldEvidence"`
				RealWorldEvidenceLibre struct {
					PolicyAccept int  `json:"policyAccept"`
					Declined     bool `json:"declined"`
					History      []struct {
						PolicyAccept int  `json:"policyAccept"`
						Declined     bool `json:"declined"`
					} `json:"history"`
				} `json:"realWorldEvidenceLibre"`
			} `json:"consents"`
			QuickShare struct {
				Code      string `json:"code"`
				ExpiresAt int    `json:"expiresAt"`
			} `json:"quickShare"`
		} `json:"user"`
		Messages struct {
			Unread int `json:"unread"`
		} `json:"messages"`
		Notifications struct {
			Unresolved int `json:"unresolved"`
		} `json:"notifications"`
		AuthTicket struct {
			Token    string `json:"token"`
			Expires  int    `json:"expires"`
			Duration int64  `json:"duration"`
		} `json:"authTicket"`
		Invitations        any    `json:"invitations"`
		TrustedDeviceToken string `json:"trustedDeviceToken"`
	} `json:"data"`
}
