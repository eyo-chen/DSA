package main

import (
	"container/heap"
)

// Tweet represents a single tweet with its ID and timestamp
type Tweet struct {
	tweetID   int // Unique identifier for the tweet
	timestamp int // Global timestamp for chronological ordering
}

// Twitter is a simplified Twitter system supporting posting tweets,
// following/unfollowing users, and retrieving news feeds
type Twitter struct {
	userTweets      map[int][]Tweet      // Maps userID to their chronologically ordered tweets
	userFollowees   map[int]map[int]bool // Maps userID to set of users they follow
	globalTimestamp int                  // Global counter for tweet ordering
}

// Constructor initializes a new Twitter system with empty data structures
// Time: O(1) - just allocating empty maps
// Space: O(1) - constant overhead for empty structures
func Constructor() Twitter {
	return Twitter{
		userTweets:      make(map[int][]Tweet),
		userFollowees:   make(map[int]map[int]bool),
		globalTimestamp: 0,
	}
}

// TweetCursor represents a pointer to a specific tweet in a user's timeline
// Used to track position when merging multiple users' tweets in the news feed
type TweetCursor struct {
	tweet  Tweet // The actual tweet data
	userID int   // ID of the user who posted this tweet
	index  int   // Position in the user's tweet array (newest tweets have highest index)
}

// TweetMaxHeap is a max heap that orders tweets by timestamp (most recent first)
// Implements heap.Interface for efficient K-way merge of sorted tweet lists
type TweetMaxHeap []TweetCursor

// Len returns the number of elements in the heap
func (h TweetMaxHeap) Len() int {
	return len(h)
}

// Less returns true if tweet at position i is MORE recent than tweet at position j
// This creates max heap behavior (highest timestamp at root)
func (h TweetMaxHeap) Less(i, j int) bool {
	return h[i].tweet.timestamp > h[j].tweet.timestamp
}

// Swap exchanges the elements at positions i and j
func (h TweetMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a new element to the heap
func (h *TweetMaxHeap) Push(x any) {
	*h = append(*h, x.(TweetCursor))
}

// Pop removes and returns the maximum element (most recent tweet) from the heap
func (h *TweetMaxHeap) Pop() any {
	oldHeap := *h
	heapSize := len(oldHeap)
	lastElement := oldHeap[heapSize-1]
	*h = oldHeap[0 : heapSize-1]
	return lastElement
}

// PostTweet creates a new tweet by the given user
// Approach: Append tweet to user's list with current timestamp, increment global counter
// Time: O(1) - amortized constant time for slice append operation
// Space: O(1) - stores one tweet object
func (t *Twitter) PostTweet(userId int, tweetId int) {
	// Create new tweet with current global timestamp
	newTweet := Tweet{
		tweetID:   tweetId,
		timestamp: t.globalTimestamp,
	}

	// Append to user's tweet list (tweets are naturally ordered by time)
	t.userTweets[userId] = append(t.userTweets[userId], newTweet)

	// Increment global timestamp for next tweet
	t.globalTimestamp++
}

// GetNewsFeed retrieves the 10 most recent tweets from the user and their followees
// Approach: K-way merge using max heap - start with most recent tweet from each user,
// extract max, then lazily load next tweet from same user. This avoids
// loading all tweets upfront and efficiently merges sorted lists.
//
// Time: O(F + K*log(F)) where F = number of followees + 1 (self), K = result size (10)
//   - O(F) to initialize heap with one tweet per user
//   - O(K*log(F)) for K iterations of pop and push operations
//
// Space: O(F) for the heap storing at most F+1 cursors, O(K) for the result array
func (t *Twitter) GetNewsFeed(userId int) []int {
	// Initialize max heap for merging tweets from multiple users
	tweetHeap := &TweetMaxHeap{}
	heap.Init(tweetHeap)

	// Add the most recent tweet from the user themselves
	if userTweetList, exists := t.userTweets[userId]; exists && len(userTweetList) > 0 {
		mostRecentIndex := len(userTweetList) - 1
		heap.Push(tweetHeap, TweetCursor{
			tweet:  userTweetList[mostRecentIndex],
			userID: userId,
			index:  mostRecentIndex,
		})
	}

	// Add the most recent tweet from each user that the current user follows
	if followeeSet, exists := t.userFollowees[userId]; exists {
		for followeeID := range followeeSet {
			if followeeTweetList, exists := t.userTweets[followeeID]; exists && len(followeeTweetList) > 0 {
				mostRecentIndex := len(followeeTweetList) - 1
				heap.Push(tweetHeap, TweetCursor{
					tweet:  followeeTweetList[mostRecentIndex],
					userID: followeeID,
					index:  mostRecentIndex,
				})
			}
		}
	}

	// Extract up to 10 most recent tweets using K-way merge algorithm
	newsFeed := make([]int, 0, 10)
	for len(newsFeed) < 10 && tweetHeap.Len() > 0 {
		// Pop the most recent tweet across all users from the heap
		mostRecentCursor := heap.Pop(tweetHeap).(TweetCursor)
		newsFeed = append(newsFeed, mostRecentCursor.tweet.tweetID)

		// Lazily load the next older tweet from the same user
		// This maintains the K-way merge without loading all tweets upfront
		if mostRecentCursor.index > 0 {
			nextOlderIndex := mostRecentCursor.index - 1
			nextTweet := t.userTweets[mostRecentCursor.userID][nextOlderIndex]
			heap.Push(tweetHeap, TweetCursor{
				tweet:  nextTweet,
				userID: mostRecentCursor.userID,
				index:  nextOlderIndex,
			})
		}
	}

	return newsFeed
}

// Follow makes followerID start following followeeID
// Approach: Add followee to follower's set using nested map structure
// Time: O(1) - hash map lookup and insertion
// Space: O(1) - stores one boolean entry in the map
func (t *Twitter) Follow(followerId int, followeeId int) {
	// Users cannot follow themselves (prevent self-follow)
	if followerId == followeeId {
		return
	}

	// Initialize follow map for this user if it doesn't exist (lazy initialization)
	if t.userFollowees[followerId] == nil {
		t.userFollowees[followerId] = make(map[int]bool)
	}

	// Add followee to follower's set (map acts as a set, duplicates are handled automatically)
	t.userFollowees[followerId][followeeId] = true
}

// Unfollow makes followerID stop following followeeID
// Approach: Remove followee from follower's set using map delete
// Time: O(1) - hash map lookup and deletion
// Space: O(1) - no additional space used
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	// Check if follower has any followees at all
	if followeeSet, exists := t.userFollowees[followerId]; exists {
		// Remove followee from the follower's set
		// Safe to delete even if followee doesn't exist in the set
		delete(followeeSet, followeeId)
	}
}
