package goinsta

import (
	"encoding/json"
	"fmt"
)

type Posts struct {
	User struct {
		EdgeOwnerToTimelineMedia struct {
			Count    int `json:"count"`
			PageInfo struct {
				HasNextPage bool   `json:"has_next_page"`
				EndCursor   string `json:"end_cursor"`
			} `json:"page_info"`
			Edges []struct {
				Node struct {
					Typename                string      `json:"__typename"`
					ID                      string      `json:"id"`
					GatingInfo              interface{} `json:"gating_info"`
					FactCheckOverallRating  interface{} `json:"fact_check_overall_rating"`
					FactCheckInformation    interface{} `json:"fact_check_information"`
					MediaOverlayInfo        interface{} `json:"media_overlay_info"`
					SensitivityFrictionInfo interface{} `json:"sensitivity_friction_info"`
					SharingFrictionInfo     struct {
						ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
						BloksAppURL               interface{} `json:"bloks_app_url"`
					} `json:"sharing_friction_info"`
					Dimensions struct {
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"dimensions"`
					DisplayURL       string `json:"display_url"`
					DisplayResources []struct {
						Src          string `json:"src"`
						ConfigWidth  int    `json:"config_width"`
						ConfigHeight int    `json:"config_height"`
					} `json:"display_resources"`
					IsVideo               bool        `json:"is_video"`
					MediaPreview          interface{} `json:"media_preview"`
					TrackingToken         string      `json:"tracking_token"`
					HasUpcomingEvent      bool        `json:"has_upcoming_event"`
					EdgeMediaToTaggedUser struct {
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_to_tagged_user"`
					AccessibilityCaption string `json:"accessibility_caption"`
					EdgeMediaToCaption   struct {
						Edges []struct {
							Node struct {
								Text string `json:"text"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_media_to_caption"`
					Shortcode          string `json:"shortcode"`
					EdgeMediaToComment struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
					} `json:"edge_media_to_comment"`
					EdgeMediaToSponsorUser struct {
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_to_sponsor_user"`
					IsAffiliate          bool `json:"is_affiliate"`
					IsPaidPartnership    bool `json:"is_paid_partnership"`
					CommentsDisabled     bool `json:"comments_disabled"`
					TakenAtTimestamp     int  `json:"taken_at_timestamp"`
					EdgeMediaPreviewLike struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_preview_like"`
					Owner struct {
						ID       string `json:"id"`
						Username string `json:"username"`
					} `json:"owner"`
					Location struct {
						ID            string `json:"id"`
						HasPublicPage bool   `json:"has_public_page"`
						Name          string `json:"name"`
						Slug          string `json:"slug"`
					} `json:"location"`
					ViewerHasLiked             bool   `json:"viewer_has_liked"`
					ViewerHasSaved             bool   `json:"viewer_has_saved"`
					ViewerHasSavedToCollection bool   `json:"viewer_has_saved_to_collection"`
					ViewerInPhotoOfYou         bool   `json:"viewer_in_photo_of_you"`
					ViewerCanReshare           bool   `json:"viewer_can_reshare"`
					ThumbnailSrc               string `json:"thumbnail_src"`
					ThumbnailResources         []struct {
						Src          string `json:"src"`
						ConfigWidth  int    `json:"config_width"`
						ConfigHeight int    `json:"config_height"`
					} `json:"thumbnail_resources"`
					CoauthorProducers     []interface{} `json:"coauthor_producers"`
					EdgeSidecarToChildren struct {
						Edges []struct {
							Node struct {
								Typename                string      `json:"__typename"`
								ID                      string      `json:"id"`
								GatingInfo              interface{} `json:"gating_info"`
								FactCheckOverallRating  interface{} `json:"fact_check_overall_rating"`
								FactCheckInformation    interface{} `json:"fact_check_information"`
								MediaOverlayInfo        interface{} `json:"media_overlay_info"`
								SensitivityFrictionInfo interface{} `json:"sensitivity_friction_info"`
								SharingFrictionInfo     struct {
									ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
									BloksAppURL               interface{} `json:"bloks_app_url"`
								} `json:"sharing_friction_info"`
								Dimensions struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL       string `json:"display_url"`
								DisplayResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"display_resources"`
								IsVideo               bool   `json:"is_video"`
								MediaPreview          string `json:"media_preview"`
								TrackingToken         string `json:"tracking_token"`
								HasUpcomingEvent      bool   `json:"has_upcoming_event"`
								EdgeMediaToTaggedUser struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
								AccessibilityCaption string `json:"accessibility_caption"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_sidecar_to_children"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_owner_to_timeline_media"`
	} `json:"user"`
}

type postResp struct {
	Data   Posts  `json:"data"`
	Status string `json:"status"`
}

type FetchPostOptions struct {
	Limit  int
	Offset string
}

func (user *User) FetchPost(opts FetchPostOptions) (*Posts, error) {
	insta := user.insta
	id := fmt.Sprintf("%d", user.ID)
	variables := map[string]string{
		"after": "",
		"first": "12",
		"id":    id,
	}

	if opts.Limit > 0 {
		variables["first"] = fmt.Sprintf("%d", opts.Limit)

	}
	if opts.Offset != "" {
		variables["after"] = opts.Offset
	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"query_hash": graphQLUserFeedHash,
		"variables":  string(variablesJSON),
	}

	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlGraphQL,
			Query:    data,
			OmitAPI:  true,
		},
	)

	resp := postResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
