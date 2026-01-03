package main

// MyListNode represents a single tweet node in a linked list for efficient feed management
type MyListNode struct {
	TweetID int         // Unique identifier for the tweet
	UserID  int         // ID of the user who posted this tweet
	Counter int         // Global timestamp for chronological ordering
	Next    *MyListNode // Pointer to the next tweet node in the list
}

// Twitter implements a social media platform using linked lists for optimized feed retrieval
type Twitter struct {
	Users     map[int]*MyListNode  // Maps userID to head of their posts list (chronologically ordered)
	Feeds     map[int]*MyListNode  // Maps userID to head of their personalized feed list
	Followers map[int]map[int]bool // Maps followeeID to set of followerIDs (reverse mapping)
	Counter   int                  // Global monotonic counter for tweet ordering
}

// Constructor initializes a new Twitter instance with empty data structures
// Time: O(1), Space: O(1)
func Constructor() Twitter {
	return Twitter{
		Users:     map[int]*MyListNode{},
		Feeds:     map[int]*MyListNode{},
		Followers: map[int]map[int]bool{},
	}
}

// PostTweet creates a new tweet and updates relevant feeds
// Approach: Insert tweet at head of user's post list and all follower feeds
// Time: O(F) where F is the number of followers the user has
// Space: O(F) for creating new nodes for each follower's feed
func (t *Twitter) PostTweet(userId int, tweetId int) {
	// Add tweet to the user's own posts list
	t.insertNodeAtHead(userId, tweetId, t.Users, userId)

	// Add tweet to the user's own feed
	t.insertNodeAtHead(userId, tweetId, t.Feeds, userId)

	// Increment global counter for next tweet
	t.Counter++

	// Add this tweet to all followers' feeds for real-time updates
	followerSet := t.Followers[userId]
	for followerID := range followerSet {
		t.insertNodeAtHead(userId, tweetId, t.Feeds, followerID)
	}
}

// insertNodeAtHead creates a new tweet node and inserts it at the head of a linked list
// Time: O(1) - constant time insertion at head
// Space: O(1) - creates one new node
func (t *Twitter) insertNodeAtHead(posterUserID int, tweetID int, listMap map[int]*MyListNode, targetUserID int) {
	// Create new node with tweet information and current timestamp
	newNode := &MyListNode{
		TweetID: tweetID,
		UserID:  posterUserID,
		Counter: t.Counter,
	}

	// Link the new node to existing list head (if exists)
	if currentHead, exists := listMap[targetUserID]; exists {
		newNode.Next = currentHead
	}

	// Update the list head to point to the new node
	listMap[targetUserID] = newNode
}

// GetNewsFeed retrieves the 10 most recent tweets from the user's personalized feed
// Approach: Traverse the pre-built feed linked list up to 10 nodes
// Time: O(1) - always retrieves at most 10 tweets
// Space: O(1) - result slice has max size 10
func (t *Twitter) GetNewsFeed(userId int) []int {
	// Pre-allocate result slice with capacity 10
	tweetIDs := make([]int, 0, 10)

	// Start from the head of user's feed (most recent tweets)
	currentNode := t.Feeds[userId]

	// Traverse up to 10 nodes in the feed
	for range 10 {
		// Stop if we've reached the end of the list
		if currentNode == nil {
			break
		}

		// Add tweet ID to result
		tweetIDs = append(tweetIDs, currentNode.TweetID)

		// Move to next tweet in the feed
		currentNode = currentNode.Next
	}

	return tweetIDs
}

// Follow creates a follow relationship and merges followee's posts into follower's feed
// Approach: Merge follower's current feed with followee's posts list while maintaining order
// Time: O(N + M) where N is follower's feed size and M is followee's posts count
// Space: O(N + M) for creating new merged feed with copied nodes
func (t *Twitter) Follow(followerId int, followeeId int) {
	// Prevent self-following
	if followerId == followeeId {
		return
	}

	// Check if already following to avoid duplicate work
	if t.Followers[followeeId] != nil {
		if _, alreadyFollowing := t.Followers[followeeId][followerId]; alreadyFollowing {
			return
		}
	}

	// Get current feeds for both users
	followerFeed := t.Feeds[followerId]
	followeePosts := t.Users[followeeId]

	// Merge follower's existing feed with followee's posts in chronological order
	mergedFeed := t.mergeFeedsWithCopy(followerFeed, followeePosts)

	// Update follower's feed with the merged result
	t.Feeds[followerId] = mergedFeed

	// Record the follow relationship in the followers map
	t.addFollowRelationship(followerId, followeeId)
}

// addFollowRelationship records a follow relationship in the followers map
// Time: O(1) - hash map operations
// Space: O(1) - stores one boolean entry
func (t *Twitter) addFollowRelationship(followerID int, followeeID int) {
	// Initialize the followee's follower set if it doesn't exist
	if t.Followers[followeeID] == nil {
		t.Followers[followeeID] = map[int]bool{}
	}

	// Add follower to the followee's follower set
	t.Followers[followeeID][followerID] = true
}

// mergeFeedsWithCopy merges two chronologically ordered linked lists by copying nodes
// Approach: Two-pointer merge similar to merge sort, creating new nodes to avoid aliasing
// Time: O(N + M) where N and M are lengths of the two lists
// Space: O(N + M) for creating copied nodes in the merged list
func (t *Twitter) mergeFeedsWithCopy(feed1 *MyListNode, feed2 *MyListNode) *MyListNode {
	// Create dummy head for easier list construction
	dummyHead := &MyListNode{}
	current := dummyHead

	// Merge both lists while both have nodes remaining
	for feed1 != nil && feed2 != nil {
		if feed1.Counter > feed2.Counter {
			// feed1 node is more recent, copy it to merged list
			current.Next = &MyListNode{
				TweetID: feed1.TweetID,
				UserID:  feed1.UserID,
				Counter: feed1.Counter,
			}
			feed1 = feed1.Next
		} else {
			// feed2 node is more recent or equal, copy it to merged list
			current.Next = &MyListNode{
				TweetID: feed2.TweetID,
				UserID:  feed2.UserID,
				Counter: feed2.Counter,
			}
			feed2 = feed2.Next
		}
		current = current.Next
	}

	// Copy remaining nodes from feed1 if any
	for feed1 != nil {
		current.Next = &MyListNode{
			TweetID: feed1.TweetID,
			UserID:  feed1.UserID,
			Counter: feed1.Counter,
		}
		current = current.Next
		feed1 = feed1.Next
	}

	// Copy remaining nodes from feed2 if any
	for feed2 != nil {
		current.Next = &MyListNode{
			TweetID: feed2.TweetID,
			UserID:  feed2.UserID,
			Counter: feed2.Counter,
		}
		current = current.Next
		feed2 = feed2.Next
	}

	// Return the merged list (skip dummy head)
	return dummyHead.Next
}

// Unfollow removes a follow relationship and deletes followee's tweets from follower's feed
// Approach: Traverse follower's feed and remove all nodes posted by followee
// Time: O(N) where N is the length of follower's feed
// Space: O(1) - only modifies existing list structure
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	// Get follower's current feed
	followerFeed := t.Feeds[followerId]

	// Use dummy head to simplify deletion logic
	dummyHead := &MyListNode{Next: followerFeed}
	current := dummyHead

	// Traverse the feed and remove all tweets from followee
	for current != nil && current.Next != nil {
		if current.Next.UserID == followeeId {
			// Skip the node (delete it from the list)
			current.Next = current.Next.Next
		} else {
			// Move to next node
			current = current.Next
		}
	}

	// Update follower's feed with modified list
	t.Feeds[followerId] = dummyHead.Next

	// Remove the follow relationship from followers map
	if t.Followers[followeeId] != nil {
		delete(t.Followers[followeeId], followerId)
	}
}
