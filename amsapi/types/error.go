package types

// ErrorType represents both the type and subtype errors.
type ErrorType string

const ErrorAccountProblem ErrorType = "ACCOUNT_PROBLEM"
const ErrorExpiredAuthorizationCredential ErrorType = "EXPIRED_AUTHORIZATION_CREDENTIAL"
const ErrorInternalError ErrorType = "INTERNAL_ERROR"
const ErrorInvalidAuthorizationCredential ErrorType = "INVALID_AUTHORIZATION_CREDENTIAL"
const ErrorInvalidDirective ErrorType = "INVALID_DIRECTIVE"
const ErrorRateLimitExceeded ErrorType = "RATE_LIMIT_EXCEEDED"

// AccountProblemInactiveSubscription subtype of ErrorAccountProblem
const AccountProblemInactiveSubscription ErrorType = "INACTIVE_SUBSCRIPTION"

// AccountProblemNotSubscribed subtype of ErrorAccountProblem
const AccountProblemNotSubscribed ErrorType = "NOT_SUBSCRIBED"

// AccountProblemPaymentProblem subtype of ErrorAccountProblem
const AccountProblemPaymentProblem ErrorType = "PAYMENT_PROBLEM"

const MediaContentFiltered ErrorType = "CONTENT_FILTERED"
const MediaInsufficientUserSubscription ErrorType = "INSUFFICIENT_USER_SUBSCRIPTION"
const MediaInvalidItem ErrorType = "INVALID_ITEM"
const MediaItemNotFound ErrorType = "ITEM_NOT_FOUND"
const MediaCriteriaNotFound ErrorType = "CRITERIA_NOT_FOUND"
const MediaContentNotFound ErrorType = "CONTENT_NOT_FOUND"
const MediaGeographicalRestrictionError ErrorType = "GEOGRAPHICAL_RESTRICTION_ERROR"
const MediaDeviceLimitReached ErrorType = "DEVICE_LIMIT_REACHED"

// MediaContentFilteredExplicit subtype of MediaContentFiltered
const MediaContentFilteredExplicit ErrorType = "EXPLICIT_LANGUAGE_FILTER"

// MediaContentNotFoundNoFollowed subtype of MediaCriteriaNotFound
const MediaContentNotFoundNoFollowed ErrorType = "NO_FOLLOWED_CONTENT"

const AudioSkipLimitReached ErrorType = "SKIP_LIMIT_REACHED"

type AudioSkipRetryPeriod string

const AudioSkipRetryPeriodHourly AudioSkipRetryPeriod = "HOURLY"
const AudioSkipRetryPeriodDaily AudioSkipRetryPeriod = "DAILY"
const AudioSkipRetryPeriodUnknown AudioSkipRetryPeriod = "UNKNOWN"
