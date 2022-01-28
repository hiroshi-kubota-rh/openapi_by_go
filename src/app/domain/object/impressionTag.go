package object

type (
	impressionTagID = uint64

	// Account account
	ImpressionTag struct {
		// The internal ID of the account
		ID impressionTagID `json:"-" db:"id"`

		// The username of the account
		Name string `json:"name" db:"name"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
