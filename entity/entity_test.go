package entity

import (
	"reflect"
	"testing"
)

var mp = &menfessPost{
	id:        "id",
	timestamp: 1000,
	name:      "name",
	avatar:    "avatar",
	body:      "body",
	upvoterIDs: map[string]bool{
		"1": true,
		"2": false,
	},
	downvoterIDs: map[string]bool{
		"3": true,
		"4": false,
	},
	replyCount: 1000,
	parentID:   "parentID",
	repostID:   "repostID",
}

var mp2 = &menfessPost{
	id:        "id2",
	timestamp: 2000,
	name:      "name2",
	avatar:    "avatar2",
	body:      "body2",
	upvoterIDs: map[string]bool{
		"11": true,
		"22": false,
	},
	downvoterIDs: map[string]bool{
		"33": true,
		"44": false,
	},
	replyCount: 2000,
	parentID:   "parentID2",
	repostID:   "repostID2",
}

func TestMenfessPostConstructor_New(t *testing.T) {
	tests := []struct {
		name string
		c    MenfessPostConstructor
		want MenfessPost
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MenfessPostConstructor.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_RepostID(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"post 1", mp, mp.repostID},
		{"post 2", mp2, mp2.repostID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.RepostID(); got != tt.want {
				t.Errorf("menfessPost.RepostID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Upvote(t *testing.T) {
	type args struct {
		accountID string
	}
	tests := []struct {
		name string
		mp   *menfessPost
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Upvote(tt.args.accountID); got != tt.want {
				t.Errorf("menfessPost.Upvote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Downvote(t *testing.T) {
	type args struct {
		accountID string
	}
	tests := []struct {
		name string
		mp   *menfessPost
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Downvote(tt.args.accountID); got != tt.want {
				t.Errorf("menfessPost.Downvote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_IsUpvoted(t *testing.T) {
	type args struct {
		accountID string
	}
	tests := []struct {
		name string
		mp   *menfessPost
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"mp 1", mp, args{"1"}, true},
		{"mp 2", mp, args{"2"}, true},
		{"mp 3", mp, args{"3"}, false},
		{"mp2 11", mp2, args{"11"}, true},
		{"mp2 33", mp2, args{"33"}, false},
		{"mp2 44", mp2, args{"44"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.IsUpvoted(tt.args.accountID); got != tt.want {
				t.Errorf("menfessPost.IsUpvoted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_IsDownvoted(t *testing.T) {
	type args struct {
		accountID string
	}
	tests := []struct {
		name string
		mp   *menfessPost
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"mp 3", mp, args{"3"}, true},
		{"mp 4", mp, args{"4"}, true},
		{"mp 1", mp, args{"1"}, false},
		{"mp2 11", mp2, args{"11"}, false},
		{"mp2 33", mp2, args{"33"}, true},
		{"mp2 44", mp2, args{"44"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.IsDownvoted(tt.args.accountID); got != tt.want {
				t.Errorf("menfessPost.IsDownvoted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_ParentID(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"post 1", mp, mp.parentID},
		{"post 2", mp2, mp2.parentID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.ParentID(); got != tt.want {
				t.Errorf("menfessPost.ParentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_ReplyCount(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want int
	}{
		// TODO: Add test cases.
		{"post 1", mp, mp.replyCount},
		{"post 2", mp2, mp2.replyCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.ReplyCount(); got != tt.want {
				t.Errorf("menfessPost.ReplyCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_DownvoteCount(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want int
	}{
		// TODO: Add test cases.
		{"post 1", mp, len(mp.downvoterIDs)},
		{"post 2", mp2, len(mp2.downvoterIDs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.DownvoteCount(); got != tt.want {
				t.Errorf("menfessPost.DownvoteCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_UpvoteCount(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want int
	}{
		// TODO: Add test cases.
		{"post 1", mp, len(mp.upvoterIDs)},
		{"post 2", mp2, len(mp2.upvoterIDs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.UpvoteCount(); got != tt.want {
				t.Errorf("menfessPost.UpvoteCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_UpvoterIDs(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want map[string]bool
	}{
		// TODO: Add test cases.
		{"upvoterIDs post 1", mp, mp.upvoterIDs},
		{"upvoterIDs post 2", mp2, mp2.upvoterIDs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.UpvoterIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menfessPost.UpvoterIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_DownvoterIDs(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want map[string]bool
	}{
		// TODO: Add test cases.
		{"downvoterIDs post 1", mp, mp.downvoterIDs},
		{"downvoterIDs post 2", mp2, mp2.downvoterIDs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.DownvoterIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menfessPost.DownvoterIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Body(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"body post 1", mp, mp.body},
		{"body post 2", mp2, mp2.body},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Body(); got != tt.want {
				t.Errorf("menfessPost.Body() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Avatar(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"avatar post 1", mp, mp.avatar},
		{"avatar post 2", mp2, mp2.avatar},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Avatar(); got != tt.want {
				t.Errorf("menfessPost.Avatar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Name(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"name post 1", mp, mp.name},
		{"name post 2", mp2, mp2.name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Name(); got != tt.want {
				t.Errorf("menfessPost.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_Timestamp(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want int
	}{
		// TODO: Add test cases.
		{"timestamp post 1", mp, mp.timestamp},
		{"timestamp post 2", mp2, mp2.timestamp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Timestamp(); got != tt.want {
				t.Errorf("menfessPost.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menfessPost_ID(t *testing.T) {
	tests := []struct {
		name string
		mp   *menfessPost
		want string
	}{
		// TODO: Add test cases.
		{"id post 1", mp, mp.id},
		{"id post 2", mp2, mp2.id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.ID(); got != tt.want {
				t.Errorf("menfessPost.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
