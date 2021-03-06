// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/feed_item_validation_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The possible validation errors of a feed item.
type FeedItemValidationErrorEnum_FeedItemValidationError int32

const (
	// No value has been specified.
	FeedItemValidationErrorEnum_UNSPECIFIED FeedItemValidationErrorEnum_FeedItemValidationError = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemValidationErrorEnum_UNKNOWN FeedItemValidationErrorEnum_FeedItemValidationError = 1
	// String is too short.
	FeedItemValidationErrorEnum_STRING_TOO_SHORT FeedItemValidationErrorEnum_FeedItemValidationError = 2
	// String is too long.
	FeedItemValidationErrorEnum_STRING_TOO_LONG FeedItemValidationErrorEnum_FeedItemValidationError = 3
	// Value is not provided.
	FeedItemValidationErrorEnum_VALUE_NOT_SPECIFIED FeedItemValidationErrorEnum_FeedItemValidationError = 4
	// Phone number format is invalid for region.
	FeedItemValidationErrorEnum_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT FeedItemValidationErrorEnum_FeedItemValidationError = 5
	// String does not represent a phone number.
	FeedItemValidationErrorEnum_INVALID_PHONE_NUMBER FeedItemValidationErrorEnum_FeedItemValidationError = 6
	// Phone number format is not compatible with country code.
	FeedItemValidationErrorEnum_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY FeedItemValidationErrorEnum_FeedItemValidationError = 7
	// Premium rate number is not allowed.
	FeedItemValidationErrorEnum_PREMIUM_RATE_NUMBER_NOT_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 8
	// Phone number type is not allowed.
	FeedItemValidationErrorEnum_DISALLOWED_NUMBER_TYPE FeedItemValidationErrorEnum_FeedItemValidationError = 9
	// Specified value is outside of the valid range.
	FeedItemValidationErrorEnum_VALUE_OUT_OF_RANGE FeedItemValidationErrorEnum_FeedItemValidationError = 10
	// Call tracking is not supported in the selected country.
	FeedItemValidationErrorEnum_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY FeedItemValidationErrorEnum_FeedItemValidationError = 11
	// Customer is not whitelisted for call tracking.
	FeedItemValidationErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING FeedItemValidationErrorEnum_FeedItemValidationError = 12
	// Country code is invalid.
	FeedItemValidationErrorEnum_INVALID_COUNTRY_CODE FeedItemValidationErrorEnum_FeedItemValidationError = 13
	// The specified mobile app id is invalid.
	FeedItemValidationErrorEnum_INVALID_APP_ID FeedItemValidationErrorEnum_FeedItemValidationError = 14
	// Some required field attributes are missing.
	FeedItemValidationErrorEnum_MISSING_ATTRIBUTES_FOR_FIELDS FeedItemValidationErrorEnum_FeedItemValidationError = 15
	// Invalid email button type for email extension.
	FeedItemValidationErrorEnum_INVALID_TYPE_ID FeedItemValidationErrorEnum_FeedItemValidationError = 16
	// Email address is invalid.
	FeedItemValidationErrorEnum_INVALID_EMAIL_ADDRESS FeedItemValidationErrorEnum_FeedItemValidationError = 17
	// The HTTPS URL in email extension is invalid.
	FeedItemValidationErrorEnum_INVALID_HTTPS_URL FeedItemValidationErrorEnum_FeedItemValidationError = 18
	// Delivery address is missing from email extension.
	FeedItemValidationErrorEnum_MISSING_DELIVERY_ADDRESS FeedItemValidationErrorEnum_FeedItemValidationError = 19
	// FeedItem scheduling start date comes after end date.
	FeedItemValidationErrorEnum_START_DATE_AFTER_END_DATE FeedItemValidationErrorEnum_FeedItemValidationError = 20
	// FeedItem scheduling start time is missing.
	FeedItemValidationErrorEnum_MISSING_FEED_ITEM_START_TIME FeedItemValidationErrorEnum_FeedItemValidationError = 21
	// FeedItem scheduling end time is missing.
	FeedItemValidationErrorEnum_MISSING_FEED_ITEM_END_TIME FeedItemValidationErrorEnum_FeedItemValidationError = 22
	// Cannot compute system attributes on a FeedItem that has no FeedItemId.
	FeedItemValidationErrorEnum_MISSING_FEED_ITEM_ID FeedItemValidationErrorEnum_FeedItemValidationError = 23
	// Call extension vanity phone numbers are not supported.
	FeedItemValidationErrorEnum_VANITY_PHONE_NUMBER_NOT_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 24
	// Invalid review text.
	FeedItemValidationErrorEnum_INVALID_REVIEW_EXTENSION_SNIPPET FeedItemValidationErrorEnum_FeedItemValidationError = 25
	// Invalid format for numeric value in ad parameter.
	FeedItemValidationErrorEnum_INVALID_NUMBER_FORMAT FeedItemValidationErrorEnum_FeedItemValidationError = 26
	// Invalid format for date value in ad parameter.
	FeedItemValidationErrorEnum_INVALID_DATE_FORMAT FeedItemValidationErrorEnum_FeedItemValidationError = 27
	// Invalid format for price value in ad parameter.
	FeedItemValidationErrorEnum_INVALID_PRICE_FORMAT FeedItemValidationErrorEnum_FeedItemValidationError = 28
	// Unrecognized type given for value in ad parameter.
	FeedItemValidationErrorEnum_UNKNOWN_PLACEHOLDER_FIELD FeedItemValidationErrorEnum_FeedItemValidationError = 29
	// Enhanced sitelinks must have both description lines specified.
	FeedItemValidationErrorEnum_MISSING_ENHANCED_SITELINK_DESCRIPTION_LINE FeedItemValidationErrorEnum_FeedItemValidationError = 30
	// Review source is ineligible.
	FeedItemValidationErrorEnum_REVIEW_EXTENSION_SOURCE_INELIGIBLE FeedItemValidationErrorEnum_FeedItemValidationError = 31
	// Review text cannot contain hyphens or dashes.
	FeedItemValidationErrorEnum_HYPHENS_IN_REVIEW_EXTENSION_SNIPPET FeedItemValidationErrorEnum_FeedItemValidationError = 32
	// Review text cannot contain double quote characters.
	FeedItemValidationErrorEnum_DOUBLE_QUOTES_IN_REVIEW_EXTENSION_SNIPPET FeedItemValidationErrorEnum_FeedItemValidationError = 33
	// Review text cannot contain quote characters.
	FeedItemValidationErrorEnum_QUOTES_IN_REVIEW_EXTENSION_SNIPPET FeedItemValidationErrorEnum_FeedItemValidationError = 34
	// Parameters are encoded in the wrong format.
	FeedItemValidationErrorEnum_INVALID_FORM_ENCODED_PARAMS FeedItemValidationErrorEnum_FeedItemValidationError = 35
	// URL parameter name must contain only letters, numbers, underscores, and
	// dashes.
	FeedItemValidationErrorEnum_INVALID_URL_PARAMETER_NAME FeedItemValidationErrorEnum_FeedItemValidationError = 36
	// Cannot find address location.
	FeedItemValidationErrorEnum_NO_GEOCODING_RESULT FeedItemValidationErrorEnum_FeedItemValidationError = 37
	// Review extension text has source name.
	FeedItemValidationErrorEnum_SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT FeedItemValidationErrorEnum_FeedItemValidationError = 38
	// Some phone numbers can be shorter than usual. Some of these short numbers
	// are carrier-specific, and we disallow those in ad extensions because they
	// will not be available to all users.
	FeedItemValidationErrorEnum_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 39
	// Triggered when a request references a placeholder field id that does not
	// exist.
	FeedItemValidationErrorEnum_INVALID_PLACEHOLDER_FIELD_ID FeedItemValidationErrorEnum_FeedItemValidationError = 40
	// URL contains invalid ValueTrack tags or format.
	FeedItemValidationErrorEnum_INVALID_URL_TAG FeedItemValidationErrorEnum_FeedItemValidationError = 41
	// Provided list exceeds acceptable size.
	FeedItemValidationErrorEnum_LIST_TOO_LONG FeedItemValidationErrorEnum_FeedItemValidationError = 42
	// Certain combinations of attributes aren't allowed to be specified in the
	// same feed item.
	FeedItemValidationErrorEnum_INVALID_ATTRIBUTES_COMBINATION FeedItemValidationErrorEnum_FeedItemValidationError = 43
	// An attribute has the same value repeatedly.
	FeedItemValidationErrorEnum_DUPLICATE_VALUES FeedItemValidationErrorEnum_FeedItemValidationError = 44
	// Advertisers can link a conversion action with a phone number to indicate
	// that sufficiently long calls forwarded to that phone number should be
	// counted as conversions of the specified type.  This is an error message
	// indicating that the conversion action specified is invalid (e.g., the
	// conversion action does not exist within the appropriate Google Ads
	// account, or it is a type of conversion not appropriate to phone call
	// conversions).
	FeedItemValidationErrorEnum_INVALID_CALL_CONVERSION_ACTION_ID FeedItemValidationErrorEnum_FeedItemValidationError = 45
	// Tracking template requires final url to be set.
	FeedItemValidationErrorEnum_CANNOT_SET_WITHOUT_FINAL_URLS FeedItemValidationErrorEnum_FeedItemValidationError = 46
	// An app id was provided that doesn't exist in the given app store.
	FeedItemValidationErrorEnum_APP_ID_DOESNT_EXIST_IN_APP_STORE FeedItemValidationErrorEnum_FeedItemValidationError = 47
	// Invalid U2 final url.
	FeedItemValidationErrorEnum_INVALID_FINAL_URL FeedItemValidationErrorEnum_FeedItemValidationError = 48
	// Invalid U2 tracking url.
	FeedItemValidationErrorEnum_INVALID_TRACKING_URL FeedItemValidationErrorEnum_FeedItemValidationError = 49
	// Final URL should start from App download URL.
	FeedItemValidationErrorEnum_INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL FeedItemValidationErrorEnum_FeedItemValidationError = 50
	// List provided is too short.
	FeedItemValidationErrorEnum_LIST_TOO_SHORT FeedItemValidationErrorEnum_FeedItemValidationError = 51
	// User Action field has invalid value.
	FeedItemValidationErrorEnum_INVALID_USER_ACTION FeedItemValidationErrorEnum_FeedItemValidationError = 52
	// Type field has invalid value.
	FeedItemValidationErrorEnum_INVALID_TYPE_NAME FeedItemValidationErrorEnum_FeedItemValidationError = 53
	// Change status for event is invalid.
	FeedItemValidationErrorEnum_INVALID_EVENT_CHANGE_STATUS FeedItemValidationErrorEnum_FeedItemValidationError = 54
	// The header of a structured snippets extension is not one of the valid
	// headers.
	FeedItemValidationErrorEnum_INVALID_SNIPPETS_HEADER FeedItemValidationErrorEnum_FeedItemValidationError = 55
	// Android app link is not formatted correctly
	FeedItemValidationErrorEnum_INVALID_ANDROID_APP_LINK FeedItemValidationErrorEnum_FeedItemValidationError = 56
	// Phone number incompatible with call tracking for country.
	FeedItemValidationErrorEnum_NUMBER_TYPE_WITH_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY FeedItemValidationErrorEnum_FeedItemValidationError = 57
	// The input is identical to a reserved keyword
	FeedItemValidationErrorEnum_RESERVED_KEYWORD_OTHER FeedItemValidationErrorEnum_FeedItemValidationError = 58
	// Each option label in the message extension must be unique.
	FeedItemValidationErrorEnum_DUPLICATE_OPTION_LABELS FeedItemValidationErrorEnum_FeedItemValidationError = 59
	// Each option prefill in the message extension must be unique.
	FeedItemValidationErrorEnum_DUPLICATE_OPTION_PREFILLS FeedItemValidationErrorEnum_FeedItemValidationError = 60
	// In message extensions, the number of optional labels and optional
	// prefills must be the same.
	FeedItemValidationErrorEnum_UNEQUAL_LIST_LENGTHS FeedItemValidationErrorEnum_FeedItemValidationError = 61
	// All currency codes in an ad extension must be the same.
	FeedItemValidationErrorEnum_INCONSISTENT_CURRENCY_CODES FeedItemValidationErrorEnum_FeedItemValidationError = 62
	// Headers in price extension are not unique.
	FeedItemValidationErrorEnum_PRICE_EXTENSION_HAS_DUPLICATED_HEADERS FeedItemValidationErrorEnum_FeedItemValidationError = 63
	// Header and description in an item are the same.
	FeedItemValidationErrorEnum_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION FeedItemValidationErrorEnum_FeedItemValidationError = 64
	// Price extension has too few items.
	FeedItemValidationErrorEnum_PRICE_EXTENSION_HAS_TOO_FEW_ITEMS FeedItemValidationErrorEnum_FeedItemValidationError = 65
	// The given value is not supported.
	FeedItemValidationErrorEnum_UNSUPPORTED_VALUE FeedItemValidationErrorEnum_FeedItemValidationError = 66
	// Invalid final mobile url.
	FeedItemValidationErrorEnum_INVALID_FINAL_MOBILE_URL FeedItemValidationErrorEnum_FeedItemValidationError = 67
	// The given string value of Label contains invalid characters
	FeedItemValidationErrorEnum_INVALID_KEYWORDLESS_AD_RULE_LABEL FeedItemValidationErrorEnum_FeedItemValidationError = 68
	// The given URL contains value track parameters.
	FeedItemValidationErrorEnum_VALUE_TRACK_PARAMETER_NOT_SUPPORTED FeedItemValidationErrorEnum_FeedItemValidationError = 69
	// The given value is not supported in the selected language of an
	// extension.
	FeedItemValidationErrorEnum_UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE FeedItemValidationErrorEnum_FeedItemValidationError = 70
	// The iOS app link is not formatted correctly.
	FeedItemValidationErrorEnum_INVALID_IOS_APP_LINK FeedItemValidationErrorEnum_FeedItemValidationError = 71
	// iOS app link or iOS app store id is missing.
	FeedItemValidationErrorEnum_MISSING_IOS_APP_LINK_OR_IOS_APP_STORE_ID FeedItemValidationErrorEnum_FeedItemValidationError = 72
	// Promotion time is invalid.
	FeedItemValidationErrorEnum_PROMOTION_INVALID_TIME FeedItemValidationErrorEnum_FeedItemValidationError = 73
	// Both the percent off and money amount off fields are set.
	FeedItemValidationErrorEnum_PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF FeedItemValidationErrorEnum_FeedItemValidationError = 74
	// Both the promotion code and orders over amount fields are set.
	FeedItemValidationErrorEnum_PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT FeedItemValidationErrorEnum_FeedItemValidationError = 75
	// Too many decimal places are specified.
	FeedItemValidationErrorEnum_TOO_MANY_DECIMAL_PLACES_SPECIFIED FeedItemValidationErrorEnum_FeedItemValidationError = 76
	// Ad Customizers are present and not allowed.
	FeedItemValidationErrorEnum_AD_CUSTOMIZERS_NOT_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 77
	// Language code is not valid.
	FeedItemValidationErrorEnum_INVALID_LANGUAGE_CODE FeedItemValidationErrorEnum_FeedItemValidationError = 78
	// Language is not supported.
	FeedItemValidationErrorEnum_UNSUPPORTED_LANGUAGE FeedItemValidationErrorEnum_FeedItemValidationError = 79
	// IF Function is present and not allowed.
	FeedItemValidationErrorEnum_IF_FUNCTION_NOT_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 80
	// Final url suffix is not valid.
	FeedItemValidationErrorEnum_INVALID_FINAL_URL_SUFFIX FeedItemValidationErrorEnum_FeedItemValidationError = 81
	// Final url suffix contains an invalid tag.
	FeedItemValidationErrorEnum_INVALID_TAG_IN_FINAL_URL_SUFFIX FeedItemValidationErrorEnum_FeedItemValidationError = 82
	// Final url suffix is formatted incorrectly.
	FeedItemValidationErrorEnum_INVALID_FINAL_URL_SUFFIX_FORMAT FeedItemValidationErrorEnum_FeedItemValidationError = 83
	// Consent for call recording, which is required for the use of call
	// extensions, was not provided by the advertiser.
	FeedItemValidationErrorEnum_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED FeedItemValidationErrorEnum_FeedItemValidationError = 84
	// Multiple message delivery options are set.
	FeedItemValidationErrorEnum_ONLY_ONE_DELIVERY_OPTION_IS_ALLOWED FeedItemValidationErrorEnum_FeedItemValidationError = 85
	// No message delivery option is set.
	FeedItemValidationErrorEnum_NO_DELIVERY_OPTION_IS_SET FeedItemValidationErrorEnum_FeedItemValidationError = 86
	// String value of conversion reporting state field is not valid.
	FeedItemValidationErrorEnum_INVALID_CONVERSION_REPORTING_STATE FeedItemValidationErrorEnum_FeedItemValidationError = 87
	// Image size is not right.
	FeedItemValidationErrorEnum_IMAGE_SIZE_WRONG FeedItemValidationErrorEnum_FeedItemValidationError = 88
	// Email delivery is not supported in the country specified in the country
	// code field.
	FeedItemValidationErrorEnum_EMAIL_DELIVERY_NOT_AVAILABLE_IN_COUNTRY FeedItemValidationErrorEnum_FeedItemValidationError = 89
	// Auto reply is not supported in the country specified in the country code
	// field.
	FeedItemValidationErrorEnum_AUTO_REPLY_NOT_AVAILABLE_IN_COUNTRY FeedItemValidationErrorEnum_FeedItemValidationError = 90
)

var FeedItemValidationErrorEnum_FeedItemValidationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "STRING_TOO_SHORT",
	3:  "STRING_TOO_LONG",
	4:  "VALUE_NOT_SPECIFIED",
	5:  "INVALID_DOMESTIC_PHONE_NUMBER_FORMAT",
	6:  "INVALID_PHONE_NUMBER",
	7:  "PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
	8:  "PREMIUM_RATE_NUMBER_NOT_ALLOWED",
	9:  "DISALLOWED_NUMBER_TYPE",
	10: "VALUE_OUT_OF_RANGE",
	11: "CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
	12: "CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING",
	13: "INVALID_COUNTRY_CODE",
	14: "INVALID_APP_ID",
	15: "MISSING_ATTRIBUTES_FOR_FIELDS",
	16: "INVALID_TYPE_ID",
	17: "INVALID_EMAIL_ADDRESS",
	18: "INVALID_HTTPS_URL",
	19: "MISSING_DELIVERY_ADDRESS",
	20: "START_DATE_AFTER_END_DATE",
	21: "MISSING_FEED_ITEM_START_TIME",
	22: "MISSING_FEED_ITEM_END_TIME",
	23: "MISSING_FEED_ITEM_ID",
	24: "VANITY_PHONE_NUMBER_NOT_ALLOWED",
	25: "INVALID_REVIEW_EXTENSION_SNIPPET",
	26: "INVALID_NUMBER_FORMAT",
	27: "INVALID_DATE_FORMAT",
	28: "INVALID_PRICE_FORMAT",
	29: "UNKNOWN_PLACEHOLDER_FIELD",
	30: "MISSING_ENHANCED_SITELINK_DESCRIPTION_LINE",
	31: "REVIEW_EXTENSION_SOURCE_INELIGIBLE",
	32: "HYPHENS_IN_REVIEW_EXTENSION_SNIPPET",
	33: "DOUBLE_QUOTES_IN_REVIEW_EXTENSION_SNIPPET",
	34: "QUOTES_IN_REVIEW_EXTENSION_SNIPPET",
	35: "INVALID_FORM_ENCODED_PARAMS",
	36: "INVALID_URL_PARAMETER_NAME",
	37: "NO_GEOCODING_RESULT",
	38: "SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT",
	39: "CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
	40: "INVALID_PLACEHOLDER_FIELD_ID",
	41: "INVALID_URL_TAG",
	42: "LIST_TOO_LONG",
	43: "INVALID_ATTRIBUTES_COMBINATION",
	44: "DUPLICATE_VALUES",
	45: "INVALID_CALL_CONVERSION_ACTION_ID",
	46: "CANNOT_SET_WITHOUT_FINAL_URLS",
	47: "APP_ID_DOESNT_EXIST_IN_APP_STORE",
	48: "INVALID_FINAL_URL",
	49: "INVALID_TRACKING_URL",
	50: "INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL",
	51: "LIST_TOO_SHORT",
	52: "INVALID_USER_ACTION",
	53: "INVALID_TYPE_NAME",
	54: "INVALID_EVENT_CHANGE_STATUS",
	55: "INVALID_SNIPPETS_HEADER",
	56: "INVALID_ANDROID_APP_LINK",
	57: "NUMBER_TYPE_WITH_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
	58: "RESERVED_KEYWORD_OTHER",
	59: "DUPLICATE_OPTION_LABELS",
	60: "DUPLICATE_OPTION_PREFILLS",
	61: "UNEQUAL_LIST_LENGTHS",
	62: "INCONSISTENT_CURRENCY_CODES",
	63: "PRICE_EXTENSION_HAS_DUPLICATED_HEADERS",
	64: "ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION",
	65: "PRICE_EXTENSION_HAS_TOO_FEW_ITEMS",
	66: "UNSUPPORTED_VALUE",
	67: "INVALID_FINAL_MOBILE_URL",
	68: "INVALID_KEYWORDLESS_AD_RULE_LABEL",
	69: "VALUE_TRACK_PARAMETER_NOT_SUPPORTED",
	70: "UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE",
	71: "INVALID_IOS_APP_LINK",
	72: "MISSING_IOS_APP_LINK_OR_IOS_APP_STORE_ID",
	73: "PROMOTION_INVALID_TIME",
	74: "PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF",
	75: "PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT",
	76: "TOO_MANY_DECIMAL_PLACES_SPECIFIED",
	77: "AD_CUSTOMIZERS_NOT_ALLOWED",
	78: "INVALID_LANGUAGE_CODE",
	79: "UNSUPPORTED_LANGUAGE",
	80: "IF_FUNCTION_NOT_ALLOWED",
	81: "INVALID_FINAL_URL_SUFFIX",
	82: "INVALID_TAG_IN_FINAL_URL_SUFFIX",
	83: "INVALID_FINAL_URL_SUFFIX_FORMAT",
	84: "CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
	85: "ONLY_ONE_DELIVERY_OPTION_IS_ALLOWED",
	86: "NO_DELIVERY_OPTION_IS_SET",
	87: "INVALID_CONVERSION_REPORTING_STATE",
	88: "IMAGE_SIZE_WRONG",
	89: "EMAIL_DELIVERY_NOT_AVAILABLE_IN_COUNTRY",
	90: "AUTO_REPLY_NOT_AVAILABLE_IN_COUNTRY",
}
var FeedItemValidationErrorEnum_FeedItemValidationError_value = map[string]int32{
	"UNSPECIFIED":                                                0,
	"UNKNOWN":                                                    1,
	"STRING_TOO_SHORT":                                           2,
	"STRING_TOO_LONG":                                            3,
	"VALUE_NOT_SPECIFIED":                                        4,
	"INVALID_DOMESTIC_PHONE_NUMBER_FORMAT":                       5,
	"INVALID_PHONE_NUMBER":                                       6,
	"PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY":                     7,
	"PREMIUM_RATE_NUMBER_NOT_ALLOWED":                            8,
	"DISALLOWED_NUMBER_TYPE":                                     9,
	"VALUE_OUT_OF_RANGE":                                         10,
	"CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY":                     11,
	"CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING":                  12,
	"INVALID_COUNTRY_CODE":                                       13,
	"INVALID_APP_ID":                                             14,
	"MISSING_ATTRIBUTES_FOR_FIELDS":                              15,
	"INVALID_TYPE_ID":                                            16,
	"INVALID_EMAIL_ADDRESS":                                      17,
	"INVALID_HTTPS_URL":                                          18,
	"MISSING_DELIVERY_ADDRESS":                                   19,
	"START_DATE_AFTER_END_DATE":                                  20,
	"MISSING_FEED_ITEM_START_TIME":                               21,
	"MISSING_FEED_ITEM_END_TIME":                                 22,
	"MISSING_FEED_ITEM_ID":                                       23,
	"VANITY_PHONE_NUMBER_NOT_ALLOWED":                            24,
	"INVALID_REVIEW_EXTENSION_SNIPPET":                           25,
	"INVALID_NUMBER_FORMAT":                                      26,
	"INVALID_DATE_FORMAT":                                        27,
	"INVALID_PRICE_FORMAT":                                       28,
	"UNKNOWN_PLACEHOLDER_FIELD":                                  29,
	"MISSING_ENHANCED_SITELINK_DESCRIPTION_LINE":                 30,
	"REVIEW_EXTENSION_SOURCE_INELIGIBLE":                         31,
	"HYPHENS_IN_REVIEW_EXTENSION_SNIPPET":                        32,
	"DOUBLE_QUOTES_IN_REVIEW_EXTENSION_SNIPPET":                  33,
	"QUOTES_IN_REVIEW_EXTENSION_SNIPPET":                         34,
	"INVALID_FORM_ENCODED_PARAMS":                                35,
	"INVALID_URL_PARAMETER_NAME":                                 36,
	"NO_GEOCODING_RESULT":                                        37,
	"SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT":                       38,
	"CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED":                  39,
	"INVALID_PLACEHOLDER_FIELD_ID":                               40,
	"INVALID_URL_TAG":                                            41,
	"LIST_TOO_LONG":                                              42,
	"INVALID_ATTRIBUTES_COMBINATION":                             43,
	"DUPLICATE_VALUES":                                           44,
	"INVALID_CALL_CONVERSION_ACTION_ID":                          45,
	"CANNOT_SET_WITHOUT_FINAL_URLS":                              46,
	"APP_ID_DOESNT_EXIST_IN_APP_STORE":                           47,
	"INVALID_FINAL_URL":                                          48,
	"INVALID_TRACKING_URL":                                       49,
	"INVALID_FINAL_URL_FOR_APP_DOWNLOAD_URL":                     50,
	"LIST_TOO_SHORT":                                             51,
	"INVALID_USER_ACTION":                                        52,
	"INVALID_TYPE_NAME":                                          53,
	"INVALID_EVENT_CHANGE_STATUS":                                54,
	"INVALID_SNIPPETS_HEADER":                                    55,
	"INVALID_ANDROID_APP_LINK":                                   56,
	"NUMBER_TYPE_WITH_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY":    57,
	"RESERVED_KEYWORD_OTHER":                                     58,
	"DUPLICATE_OPTION_LABELS":                                    59,
	"DUPLICATE_OPTION_PREFILLS":                                  60,
	"UNEQUAL_LIST_LENGTHS":                                       61,
	"INCONSISTENT_CURRENCY_CODES":                                62,
	"PRICE_EXTENSION_HAS_DUPLICATED_HEADERS":                     63,
	"ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION":                 64,
	"PRICE_EXTENSION_HAS_TOO_FEW_ITEMS":                          65,
	"UNSUPPORTED_VALUE":                                          66,
	"INVALID_FINAL_MOBILE_URL":                                   67,
	"INVALID_KEYWORDLESS_AD_RULE_LABEL":                          68,
	"VALUE_TRACK_PARAMETER_NOT_SUPPORTED":                        69,
	"UNSUPPORTED_VALUE_IN_SELECTED_LANGUAGE":                     70,
	"INVALID_IOS_APP_LINK":                                       71,
	"MISSING_IOS_APP_LINK_OR_IOS_APP_STORE_ID":                   72,
	"PROMOTION_INVALID_TIME":                                     73,
	"PROMOTION_CANNOT_SET_PERCENT_OFF_AND_MONEY_AMOUNT_OFF":      74,
	"PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT": 75,
	"TOO_MANY_DECIMAL_PLACES_SPECIFIED":                          76,
	"AD_CUSTOMIZERS_NOT_ALLOWED":                                 77,
	"INVALID_LANGUAGE_CODE":                                      78,
	"UNSUPPORTED_LANGUAGE":                                       79,
	"IF_FUNCTION_NOT_ALLOWED":                                    80,
	"INVALID_FINAL_URL_SUFFIX":                                   81,
	"INVALID_TAG_IN_FINAL_URL_SUFFIX":                            82,
	"INVALID_FINAL_URL_SUFFIX_FORMAT":                            83,
	"CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED":               84,
	"ONLY_ONE_DELIVERY_OPTION_IS_ALLOWED":                        85,
	"NO_DELIVERY_OPTION_IS_SET":                                  86,
	"INVALID_CONVERSION_REPORTING_STATE":                         87,
	"IMAGE_SIZE_WRONG":                                           88,
	"EMAIL_DELIVERY_NOT_AVAILABLE_IN_COUNTRY":                    89,
	"AUTO_REPLY_NOT_AVAILABLE_IN_COUNTRY":                        90,
}

func (x FeedItemValidationErrorEnum_FeedItemValidationError) String() string {
	return proto.EnumName(FeedItemValidationErrorEnum_FeedItemValidationError_name, int32(x))
}
func (FeedItemValidationErrorEnum_FeedItemValidationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_validation_error_3f3d939bbe5c098b, []int{0, 0}
}

// Container for enum describing possible validation errors of a feed item.
type FeedItemValidationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemValidationErrorEnum) Reset()         { *m = FeedItemValidationErrorEnum{} }
func (m *FeedItemValidationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemValidationErrorEnum) ProtoMessage()    {}
func (*FeedItemValidationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_item_validation_error_3f3d939bbe5c098b, []int{0}
}
func (m *FeedItemValidationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemValidationErrorEnum.Unmarshal(m, b)
}
func (m *FeedItemValidationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemValidationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *FeedItemValidationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemValidationErrorEnum.Merge(dst, src)
}
func (m *FeedItemValidationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemValidationErrorEnum.Size(m)
}
func (m *FeedItemValidationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemValidationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemValidationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedItemValidationErrorEnum)(nil), "google.ads.googleads.v0.errors.FeedItemValidationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.FeedItemValidationErrorEnum_FeedItemValidationError", FeedItemValidationErrorEnum_FeedItemValidationError_name, FeedItemValidationErrorEnum_FeedItemValidationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/feed_item_validation_error.proto", fileDescriptor_feed_item_validation_error_3f3d939bbe5c098b)
}

var fileDescriptor_feed_item_validation_error_3f3d939bbe5c098b = []byte{
	// 1577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xef, 0x72, 0x1b, 0xb7,
	0x11, 0xaf, 0x9d, 0x36, 0x69, 0xe1, 0x26, 0x86, 0xe1, 0x3f, 0x92, 0x2d, 0x59, 0xb6, 0x65, 0xc7,
	0x76, 0x1c, 0x87, 0x52, 0x9b, 0xa6, 0x69, 0xe8, 0x36, 0x29, 0x78, 0xd8, 0x23, 0x51, 0xe1, 0x80,
	0x13, 0x80, 0x23, 0x45, 0x8d, 0x66, 0x30, 0x6a, 0xc9, 0x6a, 0x34, 0x63, 0x89, 0x19, 0x51, 0xf5,
	0xa3, 0xf4, 0x01, 0xfa, 0xb1, 0x8f, 0xd2, 0x47, 0xe9, 0x03, 0xb4, 0x5f, 0x33, 0x8b, 0xfb, 0xc3,
	0x53, 0x28, 0x79, 0xf2, 0x89, 0xc7, 0xdb, 0xc5, 0x62, 0xf1, 0xdb, 0xdf, 0xfe, 0x16, 0x47, 0xbe,
	0x3b, 0x9a, 0xcd, 0x8e, 0xde, 0x4e, 0xb7, 0x0e, 0x27, 0xf3, 0xad, 0xf2, 0x11, 0x9f, 0xde, 0x6d,
	0x6f, 0x4d, 0xcf, 0xce, 0x66, 0x67, 0xf3, 0xad, 0xbf, 0x4f, 0xa7, 0x93, 0x70, 0x7c, 0x3e, 0x3d,
	0x09, 0xef, 0x0e, 0xdf, 0x1e, 0x4f, 0x0e, 0xcf, 0x8f, 0x67, 0xa7, 0x21, 0xda, 0x3a, 0xdf, 0x9f,
	0xcd, 0xce, 0x67, 0x6c, 0xa3, 0x5c, 0xd5, 0x39, 0x9c, 0xcc, 0x3b, 0x4d, 0x80, 0xce, 0xbb, 0xed,
	0x4e, 0x19, 0x60, 0xf3, 0x9f, 0xab, 0x64, 0x2d, 0x9d, 0x4e, 0x27, 0xf2, 0x7c, 0x7a, 0x32, 0x6c,
	0x42, 0x00, 0x1a, 0xe1, 0xf4, 0x1f, 0x27, 0x9b, 0xff, 0x5b, 0x21, 0x2b, 0x57, 0xd8, 0xd9, 0x4d,
	0x72, 0xa3, 0xd0, 0x2e, 0x87, 0x44, 0xa6, 0x12, 0x04, 0xfd, 0x19, 0xbb, 0x41, 0x3e, 0x2a, 0xf4,
	0x8e, 0x36, 0x23, 0x4d, 0xaf, 0xb1, 0x3b, 0x84, 0x3a, 0x6f, 0xa5, 0xee, 0x07, 0x6f, 0x4c, 0x70,
	0x03, 0x63, 0x3d, 0xbd, 0xce, 0x6e, 0x93, 0x9b, 0xad, 0xb7, 0xca, 0xe8, 0x3e, 0xfd, 0x80, 0xad,
	0x90, 0xdb, 0x43, 0xae, 0x0a, 0x08, 0xda, 0xf8, 0xb0, 0x08, 0xf8, 0x73, 0xf6, 0x92, 0x3c, 0x93,
	0x7a, 0xc8, 0x95, 0x14, 0x41, 0x98, 0x0c, 0x9c, 0x97, 0x49, 0xc8, 0x07, 0x46, 0x43, 0xd0, 0x45,
	0xd6, 0x03, 0x1b, 0x52, 0x63, 0x33, 0xee, 0xe9, 0x2f, 0xd8, 0x2a, 0xb9, 0x53, 0x7b, 0xb6, 0x1d,
	0xe8, 0x87, 0xec, 0x15, 0x79, 0x7e, 0x61, 0x49, 0xdc, 0xa3, 0xc8, 0x73, 0x63, 0x3d, 0x08, 0x0c,
	0x10, 0x12, 0x53, 0x68, 0x6f, 0xc7, 0xf4, 0x23, 0xf6, 0x94, 0x3c, 0xca, 0x2d, 0x64, 0xb2, 0xc8,
	0x82, 0xe5, 0xfe, 0xc2, 0x12, 0xae, 0x94, 0x19, 0x81, 0xa0, 0xbf, 0x64, 0x0f, 0xc8, 0x3d, 0x21,
	0x5d, 0xf5, 0xbf, 0x76, 0xf1, 0xe3, 0x1c, 0xe8, 0xaf, 0xd8, 0x3d, 0xc2, 0xca, 0x93, 0x98, 0xc2,
	0x07, 0x93, 0x06, 0xcb, 0x75, 0x1f, 0x28, 0xc1, 0x24, 0x12, 0xae, 0x94, 0xb7, 0x3c, 0xd9, 0xc1,
	0xc3, 0x5f, 0x9d, 0xc4, 0x0d, 0xf6, 0x05, 0xf9, 0x2c, 0x29, 0x9c, 0x37, 0x59, 0xb5, 0xf3, 0x68,
	0x20, 0x3d, 0x28, 0xe9, 0x1a, 0xcf, 0x56, 0x20, 0xfa, 0xeb, 0xf6, 0xc9, 0xab, 0x18, 0x21, 0x31,
	0x02, 0xe8, 0xc7, 0x8c, 0x91, 0x4f, 0x6a, 0x0b, 0xcf, 0xf3, 0x20, 0x05, 0xfd, 0x84, 0x3d, 0x21,
	0x0f, 0x33, 0xe9, 0x1c, 0xe6, 0xc0, 0xbd, 0xb7, 0xb2, 0x57, 0x78, 0x70, 0x31, 0x6c, 0x2a, 0x41,
	0x09, 0x47, 0x6f, 0x62, 0x89, 0xea, 0x65, 0x78, 0x2a, 0x5c, 0x47, 0xd9, 0x7d, 0x72, 0xb7, 0x7e,
	0x09, 0x19, 0x97, 0x2a, 0x70, 0x21, 0x2c, 0x38, 0x47, 0x6f, 0xb1, 0xbb, 0xe4, 0x56, 0x6d, 0x1a,
	0x78, 0x9f, 0xbb, 0x50, 0x58, 0x45, 0x19, 0x5b, 0x27, 0xab, 0xf5, 0x4e, 0x02, 0x94, 0x1c, 0x82,
	0x1d, 0x37, 0x8b, 0x6e, 0xb3, 0x87, 0xe4, 0xbe, 0xf3, 0xdc, 0xfa, 0x20, 0x10, 0x67, 0x9e, 0x7a,
	0xb0, 0x01, 0xb4, 0x88, 0x7f, 0xe9, 0x1d, 0xf6, 0x98, 0xac, 0xd7, 0x8b, 0x53, 0x00, 0x11, 0xa4,
	0x87, 0x2c, 0x94, 0x0b, 0xbc, 0xcc, 0x80, 0xde, 0x65, 0x1b, 0xe4, 0xc1, 0xb2, 0x07, 0x46, 0x88,
	0xf6, 0x7b, 0x08, 0xcb, 0xb2, 0x5d, 0x0a, 0xba, 0x82, 0x45, 0x1e, 0x72, 0x2d, 0xfd, 0x38, 0x2c,
	0xf1, 0xa2, 0x2e, 0xf2, 0x2a, 0x7b, 0x46, 0x1e, 0xd7, 0x87, 0xb2, 0x30, 0x94, 0x30, 0x0a, 0xb0,
	0xe7, 0x41, 0x3b, 0x69, 0x74, 0x70, 0x5a, 0xe6, 0x39, 0x78, 0x7a, 0xbf, 0x8d, 0xca, 0x45, 0x42,
	0x3e, 0x40, 0x4e, 0x37, 0xd4, 0xc5, 0x23, 0x56, 0x86, 0xb5, 0x0b, 0x4c, 0xb5, 0x32, 0x69, 0x2c,
	0xeb, 0x88, 0x49, 0xd5, 0x3e, 0x21, 0x57, 0x3c, 0x81, 0x81, 0x51, 0x02, 0xaa, 0xc2, 0xd0, 0x87,
	0xac, 0x43, 0x5e, 0xd5, 0x27, 0x02, 0x3d, 0xe0, 0x3a, 0x01, 0x11, 0x5c, 0xa4, 0x86, 0xde, 0x09,
	0x02, 0x5c, 0x62, 0x65, 0xee, 0x31, 0x3f, 0x25, 0x35, 0xd0, 0x0d, 0xf6, 0x9c, 0x6c, 0x2e, 0xa7,
	0x6e, 0x0a, 0x9b, 0x40, 0x90, 0x1a, 0x94, 0xec, 0xcb, 0x9e, 0x02, 0xfa, 0x88, 0xbd, 0x20, 0x4f,
	0x07, 0xe3, 0x7c, 0x00, 0xda, 0x05, 0xa9, 0xaf, 0x3e, 0xed, 0x63, 0x24, 0xa6, 0x30, 0x45, 0x4f,
	0x41, 0xd8, 0x2d, 0x0c, 0xd2, 0xe6, 0x7d, 0xee, 0x4f, 0x70, 0xff, 0x9f, 0xe0, 0xb7, 0xc9, 0x1e,
	0x91, 0xb5, 0x1a, 0x10, 0x84, 0x22, 0x80, 0x46, 0xfe, 0x8a, 0x90, 0x73, 0xcb, 0x33, 0x47, 0x9f,
	0x62, 0xa9, 0x6b, 0x87, 0xc2, 0xaa, 0xf2, 0x3d, 0x20, 0x61, 0x34, 0xcf, 0x80, 0x3e, 0x43, 0xa8,
	0xb5, 0x09, 0x7d, 0x30, 0x89, 0x11, 0x88, 0x8e, 0x05, 0x57, 0x28, 0x4f, 0x3f, 0x45, 0xf9, 0xa8,
	0x0e, 0x8c, 0x9e, 0x97, 0xa6, 0xe1, 0x61, 0xcf, 0xd3, 0xe7, 0xb1, 0xe7, 0xb8, 0xb5, 0x12, 0x6c,
	0xad, 0x3f, 0x49, 0x29, 0x59, 0x97, 0xb1, 0xe3, 0x05, 0xd2, 0xb3, 0xa9, 0xe1, 0x8f, 0x2b, 0x85,
	0x24, 0x7b, 0xd9, 0x6e, 0x22, 0xcc, 0xd9, 0xf3, 0x3e, 0xfd, 0x8c, 0xdd, 0x22, 0x1f, 0x63, 0x1f,
	0x2f, 0xa4, 0xef, 0x15, 0xdb, 0x24, 0x1b, 0x4d, 0x8f, 0x2e, 0xfa, 0x31, 0x31, 0x59, 0x4f, 0x6a,
	0x8e, 0xd5, 0xa4, 0x9f, 0xa3, 0x92, 0x8a, 0x22, 0x57, 0x32, 0x41, 0x1e, 0x45, 0x79, 0x71, 0xf4,
	0x35, 0xfb, 0x94, 0x3c, 0x69, 0xfa, 0x9e, 0x2b, 0x15, 0x12, 0xa3, 0x87, 0x60, 0xe3, 0xa1, 0x78,
	0x12, 0x79, 0x20, 0x05, 0xfd, 0x02, 0x1b, 0x3e, 0xe1, 0x3a, 0xea, 0x0d, 0xf8, 0x30, 0x92, 0x7e,
	0x80, 0xd2, 0x94, 0x4a, 0xcd, 0x15, 0x66, 0xe6, 0x68, 0x07, 0xb9, 0x5e, 0xea, 0x43, 0x10, 0x06,
	0x9c, 0xf6, 0x01, 0xf6, 0x30, 0x4b, 0xa9, 0xa3, 0x6c, 0x38, 0x6f, 0x2c, 0xd0, 0xad, 0x76, 0x9b,
	0x37, 0xab, 0xe9, 0x76, 0x9b, 0xce, 0x8d, 0xba, 0xa1, 0xe5, 0x37, 0xa8, 0x79, 0x4b, 0x0b, 0xa2,
	0xd2, 0x60, 0x54, 0x61, 0x46, 0x5a, 0x19, 0x1e, 0xd1, 0xa1, 0xbf, 0x45, 0xa9, 0x6a, 0x90, 0x29,
	0x47, 0xc5, 0x97, 0xed, 0x0e, 0x2a, 0x1c, 0xd8, 0xea, 0x54, 0xf4, 0x77, 0xed, 0x4c, 0xa2, 0x40,
	0x45, 0x1a, 0x7c, 0xd5, 0xe6, 0x11, 0x0c, 0x41, 0xfb, 0x90, 0x0c, 0x50, 0x7d, 0x51, 0x36, 0x7c,
	0xe1, 0xe8, 0xef, 0xd9, 0x1a, 0x59, 0xa9, 0x1d, 0x2a, 0xf6, 0xb9, 0x30, 0x00, 0x2e, 0xc0, 0xd2,
	0xaf, 0x51, 0xae, 0x9a, 0x42, 0x68, 0x61, 0x4d, 0x25, 0x9a, 0xd8, 0x5c, 0xf4, 0x0f, 0xec, 0x0d,
	0xf9, 0xba, 0x25, 0xf4, 0x11, 0xc6, 0xf0, 0x13, 0x05, 0xfd, 0x1b, 0x1c, 0x18, 0x16, 0x1c, 0xd8,
	0x21, 0x88, 0xb0, 0x03, 0xe3, 0x91, 0xb1, 0x22, 0x18, 0x3f, 0x00, 0x4b, 0xbb, 0x98, 0xd3, 0xa2,
	0xb6, 0xa6, 0xea, 0x5f, 0xde, 0x03, 0xe5, 0xe8, 0x1b, 0x14, 0x84, 0x25, 0x63, 0x6e, 0x21, 0x95,
	0x4a, 0x39, 0xfa, 0x47, 0x84, 0xbe, 0xd0, 0xb0, 0x5b, 0x70, 0x15, 0x22, 0x78, 0x0a, 0x74, 0xdf,
	0x0f, 0x1c, 0xfd, 0x53, 0x09, 0x45, 0x62, 0xb4, 0xc3, 0xa9, 0x81, 0x48, 0x14, 0xd6, 0x82, 0x4e,
	0xca, 0xc9, 0xe0, 0xe8, 0xb7, 0x71, 0x28, 0x46, 0xf1, 0x59, 0x74, 0xc2, 0x80, 0xbb, 0xd0, 0xec,
	0x26, 0x2a, 0x60, 0x1c, 0xfd, 0x0e, 0x75, 0x27, 0x8a, 0xe7, 0xa5, 0x0e, 0x08, 0x56, 0x5b, 0x7c,
	0xe8, 0x9f, 0x91, 0x98, 0x97, 0xc5, 0xc6, 0xd2, 0xa6, 0x30, 0x8a, 0x42, 0xec, 0x28, 0xc7, 0x2a,
	0x16, 0x7a, 0x01, 0x59, 0xe4, 0x35, 0xed, 0xb5, 0xeb, 0x50, 0xb2, 0x26, 0x33, 0x3d, 0xa9, 0x20,
	0xf2, 0x24, 0x69, 0x93, 0xbe, 0x42, 0x52, 0x81, 0x73, 0x81, 0x8b, 0x60, 0x0b, 0x05, 0x25, 0x72,
	0x54, 0xa0, 0xa4, 0x95, 0x63, 0x38, 0xd6, 0xa7, 0xad, 0x18, 0xed, 0x32, 0x51, 0x40, 0x1c, 0x96,
	0x92, 0x40, 0xe2, 0x3b, 0x50, 0x90, 0xe0, 0x1b, 0xc5, 0x75, 0xbf, 0xe0, 0x7d, 0xa0, 0x69, 0x9b,
	0xe9, 0xd2, 0xb8, 0x05, 0x3b, 0xfa, 0xec, 0x35, 0x79, 0x59, 0x2b, 0x73, 0xdb, 0x12, 0x8c, 0x6d,
	0xfe, 0xc7, 0x2e, 0xc2, 0x8e, 0x1c, 0x20, 0x1d, 0x72, 0x6b, 0x32, 0x53, 0xf6, 0x68, 0x4d, 0x64,
	0x9c, 0x5a, 0x92, 0x7d, 0x43, 0xbe, 0x5a, 0xd8, 0x5a, 0x7d, 0x9b, 0x83, 0x4d, 0xb0, 0x90, 0x26,
	0x4d, 0x23, 0xe2, 0x99, 0xd1, 0x30, 0x0e, 0x3c, 0x43, 0x8e, 0xe1, 0x4b, 0xfa, 0x17, 0xf6, 0x2d,
	0xe9, 0x5e, 0xbe, 0x74, 0xf1, 0xd2, 0x08, 0x88, 0xab, 0x8d, 0xc5, 0xd2, 0x06, 0x33, 0xc4, 0xfa,
	0xc5, 0x18, 0x74, 0x07, 0xa1, 0xc5, 0x12, 0x65, 0x5c, 0x8f, 0x83, 0x80, 0x44, 0x66, 0x5c, 0x95,
	0xe2, 0xe6, 0x5a, 0x57, 0x32, 0x85, 0x62, 0xcc, 0x45, 0x28, 0x2f, 0x28, 0x72, 0x1f, 0xc3, 0xb4,
	0xa5, 0x31, 0x6b, 0x8f, 0xc4, 0x1a, 0xbb, 0xf2, 0x3e, 0xa2, 0x4b, 0xbe, 0x2e, 0xc0, 0x6e, 0xa0,
	0x35, 0xb1, 0x33, 0xd3, 0x90, 0x16, 0xba, 0x14, 0xae, 0x76, 0xc4, 0x7c, 0x99, 0x11, 0xa8, 0x23,
	0xae, 0x48, 0x53, 0xb9, 0x47, 0x77, 0x71, 0x9a, 0x37, 0x18, 0xf2, 0x3e, 0xd6, 0x6e, 0xc9, 0xc9,
	0xb6, 0x9d, 0x7e, 0x6c, 0xad, 0xc7, 0xaf, 0x63, 0xdb, 0xe4, 0x75, 0x73, 0xef, 0xc2, 0xde, 0x41,
	0xbc, 0xeb, 0xfb, 0x56, 0xb0, 0x90, 0x18, 0x5b, 0x4d, 0x97, 0xdd, 0x42, 0x5a, 0x10, 0xd4, 0x23,
	0xcd, 0x8c, 0x56, 0xe3, 0x80, 0xb7, 0x88, 0xe6, 0x8e, 0x53, 0xf5, 0xa9, 0x74, 0xcd, 0x11, 0x0a,
	0x6c, 0x64, 0x6d, 0x2e, 0x73, 0x71, 0xe0, 0xe9, 0x10, 0x27, 0xe5, 0xe2, 0x0a, 0xd7, 0xa8, 0xb8,
	0x05, 0x04, 0x0a, 0xb7, 0x44, 0x05, 0x03, 0x3a, 0xc2, 0x41, 0x20, 0x33, 0x04, 0xd4, 0xc9, 0x7d,
	0x08, 0x23, 0x8b, 0x23, 0x64, 0x8f, 0x7d, 0x4e, 0x5e, 0x94, 0x57, 0xb2, 0x26, 0x7e, 0xc4, 0x6f,
	0xc8, 0xa5, 0xe2, 0x38, 0xac, 0xa5, 0x6e, 0xb4, 0x68, 0x8c, 0x29, 0xf3, 0xc2, 0x1b, 0x0c, 0xae,
	0xde, 0xe3, 0xb8, 0xdf, 0xfb, 0xff, 0x35, 0xb2, 0xf9, 0xb7, 0xd9, 0x49, 0xe7, 0xfd, 0xdf, 0x0f,
	0xbd, 0xf5, 0x2b, 0x3e, 0x0e, 0x72, 0xfc, 0xfa, 0xc8, 0xaf, 0xed, 0x8b, 0x6a, 0xfd, 0xd1, 0xec,
	0xed, 0xe1, 0xe9, 0x51, 0x67, 0x76, 0x76, 0xb4, 0x75, 0x34, 0x3d, 0x8d, 0xdf, 0x26, 0xf5, 0x07,
	0xcd, 0xf7, 0xc7, 0xf3, 0xab, 0xbe, 0x6f, 0xde, 0x94, 0x3f, 0xff, 0xba, 0xfe, 0x41, 0x9f, 0xf3,
	0x7f, 0x5f, 0xdf, 0xe8, 0x97, 0xc1, 0xf8, 0x64, 0xde, 0x29, 0x1f, 0xf1, 0x69, 0xb8, 0xdd, 0x89,
	0x5b, 0xce, 0xff, 0x53, 0x3b, 0x1c, 0xf0, 0xc9, 0xfc, 0xa0, 0x71, 0x38, 0x18, 0x6e, 0x1f, 0x94,
	0x0e, 0xff, 0xbd, 0xbe, 0x59, 0xbe, 0xed, 0x76, 0xf9, 0x64, 0xde, 0xed, 0x36, 0x2e, 0xdd, 0xee,
	0x70, 0xbb, 0xdb, 0x2d, 0x9d, 0xfe, 0xfa, 0x61, 0xcc, 0xee, 0xcb, 0x1f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x87, 0x33, 0x57, 0x7c, 0x0d, 0x00, 0x00,
}
