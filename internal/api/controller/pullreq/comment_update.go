// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pullreq

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type CommentUpdateInput struct {
	Text    string                 `json:"text"`
	Payload map[string]interface{} `json:"payload"`
}

// CommentUpdate updates a pull request comment.
func (c *Controller) CommentUpdate(
	ctx context.Context,
	session *auth.Session,
	repoRef string,
	prNum int64,
	commentID int64,
	in *CommentUpdateInput,
) (*types.PullReqActivity, error) {
	repo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoEdit)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire access to repo: %w", err)
	}

	pr, err := c.pullreqStore.FindByNumber(ctx, repo.ID, prNum)
	if err != nil {
		return nil, fmt.Errorf("failed to find pull request by number: %w", err)
	}

	act, err := c.getCommentCheckEditAccess(ctx, session, pr, commentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment: %w", err)
	}

	if act.Text == in.Text && reflect.DeepEqual(act.Payload, in.Payload) {
		return act, nil
	}

	now := time.Now().UnixMilli()
	act.Edited = now
	act.Text = in.Text
	act.Payload = in.Payload

	err = c.pullreqActivityStore.Update(ctx, act)
	if err != nil {
		return nil, fmt.Errorf("failed to update comment: %w", err)
	}

	return act, nil
}
