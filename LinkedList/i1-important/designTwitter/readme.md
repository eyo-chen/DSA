# Problem Explanation

## Naive Approach

This solution implements a simplified version of Twitter that supports **posting tweets**, **following/unfollowing users**, and **retrieving a user’s news feed**.

The design uses two main data structures:
* `userPosts`: maps each user to a list of all their tweets.
* `userFollows`: maps each user (follower) to a set of users they follow.

When a user requests their news feed, the system gathers all tweets from:

* the user themselves, and
* every user they follow.

Then it sorts these tweets by timestamp and returns the top 10 most recent ones.

This approach is straightforward and easy to implement, though `GetNewsFeed` has a higher time complexity because it collects and sorts all relevant posts.

### Component-by-Component Breakdown

#### `type post`

A lightweight struct storing metadata for each tweet:

* `userID`: author of the tweet
* `tweetID`: tweet identifier
* `timeCounter`: a global, ever-increasing timestamp to preserve ordering

#### `type Twitter`

Holds all state of the platform:

* `userPosts map[int][]post`
  * For each user → list of their tweets in chronological order.
* `userFollows map[int]map[int]bool`
  * For each follower → set of followees.
* `timeCounter int`
  * Monotonic counter incremented on each tweet to ensure stable ordering.

### Operations

#### `PostTweet(userId, tweetId)`

**What it does**
* Creates a new `post` with the current `timeCounter`.
* Appends it to the user’s list of tweets.
* Increments global `timeCounter`.

**Time Complexity O(1)**
- appending to slice is amortized O(1).

**Space Complexity O(1):**
- only storing one new post object.


#### `GetNewsFeed(userId)`

**What it does**
1. Retrieves all users that `userId` follows.
2. Collects all relevant tweets:
   * user’s own tweets
   * tweets from all followees
3. Sorts all collected tweets by `timeCounter` (descending).
4. Returns up to the 10 most recent tweet IDs.

**Time Complexity O(N log N)**
- N = total posts from user + all followees
- Most expensive step is sorting.

**Space Complexity O(N)**
- Temporary slice holding all relevant posts.

#### `Follow(followerId, followeeId)`

**What it does**
* Ensures `followerId` has an initialized followee set.
* Inserts `followeeId` into the set.

**Time Complexity O(1)**
- hash map insert.

**Space Complexity O(1)**
- adding a boolean entry.

#### `Unfollow(followerId, followeeId)`

**What it does**
* Deletes `followeeId` from the follower’s follow set.
* Safe even if the follow relationship doesn’t exist.

**Time Complexity O(1)**
- hash map delete.

**Space Complexity O(1)**
- no extra memory used.


### Summary Table
| Operation     | Description                                 | Time Complexity | Space Complexity |
| ------------- | ------------------------------------------- | --------------- | ---------------- |
| `PostTweet`   | Append one tweet to user’s list             | O(1)            | O(1)             |
| `GetNewsFeed` | Collect + sort posts, return 10 most recent | O(N log N)      | O(N)             |
| `Follow`      | Add followee to follower’s set              | O(1)            | O(1)             |
| `Unfollow`    | Remove followee from follower’s set         | O(1)            | O(1)             |


## Linked List Approach

This design uses **linked lists to maintain each user’s posts and personalized feed**, enabling constant-time insertion for new tweets and ***constant-time retrieval*** of the top 10 feed tweets.

Unlike the simpler “collect + sort” approach, this solution **pre-computes feeds**:

* When a user posts a tweet:
  * The tweet is inserted at the head of:
    * their own post list
    * their own feed list
    * every follower’s feed list
      This gives push-based, real-time feed updates.

* When a user follows someone:
  * The followee’s post list is merged with the follower’s feed.
  * The merge is similar to merging two sorted linked lists based on timestamps.

* When a user unfollows someone:
  * All tweets from that followee are removed directly from the follower’s feed.

Because feeds are always kept sorted, `GetNewsFeed()` simply returns the first 10 nodes.

The trade-off:
* **Posting becomes more expensive** (must update every follower’s feed).
* **Retrieving the feed becomes O(1)**.

### Components Breakdown

#### MyListNode

A node in the linked list representing a tweet:
* `TweetID`: ID of the tweet
* `UserID`: author of the tweet
* `Counter`: timestamp for ordering
* `Next`: pointer to next node

Used for both:
* the user's posts
* the user’s feed

#### Twitter

Stores all platform-level data:
* `Users map[int]*MyListNode`
  * Maps userID → head of user’s own posts linked list (most recent at head).
* `Feeds map[int]*MyListNode`
  * Maps userID → head of user's **personalized feed** (combined, sorted tweets).
* `Followers map[int]map[int]bool`
  * Maps **followeeID → set of followerIDs**
  * Reverse mapping makes pushing updates efficient.
* `Counter int`
  * Global monotonic timestamp incremented per tweet.


### Operations Explained

#### `PostTweet(userId, tweetId)`

**What it does**
1. Creates a tweet node.
2. Inserts it at the head of:
   * user’s own `Users[userId]` list
   * user’s own `Feeds[userId]` list
3. Pushes the tweet to **every follower’s feed**:
   * inserts at head of each follower’s feed list.
4. Increments timestamp.

**Why this works**
* The most recent tweet is always placed at the head.
* Followers' feeds are instantly updated and remain sorted.

**Time Complexity O(F)**
- F = the number of followers
- must update each follower’s feed

**Space Complexity O(1)**


#### `GetNewsFeed(userId)`

**What it does**
* Reads **at most the first 10 nodes** from the user's feed list.
* Since feeds are pre-sorted during inserts, this operation becomes trivial.

**Time Complexity O(1)**
- always read ≤10 nodes

**Space Complexity O(1)**


#### `Follow(followerId, followeeId)`

**What it does**
1. Prevents self-follow and duplicate follow.
2. Merges the following two list into a single, sorted linked list.
   * follower’s current feed
   * followee’s full post list
3. Replaces follower’s feed with the merged list.
4. Records the follow relationship.

**Merging strategy**
* Similar to merging two sorted lists (merge sort technique)
* Uses `Counter` to maintain chronological order

**Time Complexity O(N + M)**
- N = size of follower’s feed
- M = size of followee’s posts list
- we need to traverse both lists to merge them

**Space Complexity O(N + M)**
- N = size of follower’s feed
- M = size of followee’s posts list
- we need to create a new list to store the merged result

#### `Unfollow(followerId, followeeId)`

**What it does**
1. Traverses follower’s feed.
2. Removes all nodes where `UserID == followeeId`.
3. Removes the follow relationship.
3. Records the unfollow relationship.

**Time Complexity O(N)**
- N = size of follower’s feed
- we need to traverse the list to remove the nodes

**Space Complexity O(1)**
- we need to modify the list in place


### Summary Table

| Operation            | Description                          | Time         | Space        |
| -------------------- | ------------------------------------ | ------------ | ------------ |
| `PostTweet`          | Insert tweet into all relevant feeds | **O(F)**     | **O(F)**     |
| `GetNewsFeed`        | Retrieve first 10 feed nodes         | **O(1)**     | **O(1)**     |
| `Follow`             | Merge feed with followee's posts     | **O(N + M)** | **O(N + M)** |
| `Unfollow`           | Remove followee’s posts from feed    | **O(N)**     | **O(1)**     |




## Max Heap (K-way Merge) Approach

This implementation models Twitter using **arrays for storing tweets** and a **max heap to merge timelines** when retrieving the news feed.

* Each user stores their tweets in a slice (dynamic array), in chronological order.
* `GetNewsFeed()` performs a **K-way merge** across:
  * the user’s own tweets
  * all followees’ tweets
* It uses a **max heap** to always pick the most recent tweet across all timelines.
* Instead of merging all tweets upfront, it loads tweets **lazily**:
  * Only the most recent tweet of each user is pushed to the heap initially.
  * When one is popped, the next older tweet from that same user is pushed.

This keeps the operation efficient, especially when users have long tweet histories.

### Components Breakdown

#### Tweet

Represents a single tweet:
* `tweetID`: unique ID
* `timestamp`: global timestamp used for chronological ordering

#### Twitter

Stores all high-level system data:
* `userTweets map[int][]Tweet`
  * Stores each user’s tweets in an array (chronologically sorted).

* `userFollowees map[int]map[int]bool`
  * Maps each user → set of users they follow.

* `globalTimestamp`
  * Incremented per tweet to maintain strict monotonic order.

#### TweetCursor

A pointer to a specific tweet within a user’s list:
* `tweet`: the tweet itself
* `userID`: owner of the tweet
* `index`: index into `userTweets[userID]` (most recent = highest index)
This is essential for lazy loading in the K-way merge.

### Operations Explained (with Complexity)
#### `PostTweet(userId, tweetId)`

**What it does**
* Creates a new `Tweet` using the current timestamp.
* Appends to `userTweets[userId]`.
* Increments the global timestamp.

Using slice append keeps operations efficient.

**Time Complexity O(1)**

**Space Complexity O(1)**

#### `GetNewsFeed(userId)`
Uses **K-way merge with a max heap**:
1. Push the most recent tweet of the user (if any).
2. Push the most recent tweet of each followee.
3. Repeat up to 10 times:
   * Pop the most recent tweet (max timestamp).
   * Append its `tweetID` to the feed.
   * Push the **next older tweet** from that user (if exists).

This produces the top 10 most recent tweets across all timelines.

**Why this is efficient**
* Only keeps **F + 1** heap entries at any time (`F = number of followees`).
* Only processes **K=10** tweets for the news feed.
* Fully avoids scanning all tweets.


**Time Complexity O(F + K log F)**
- F = number of followees + 1 (self)
- K = max tweets to return (10)

**Space Complexity O(F)**
- F = number of followees + 1 (self)

#### `Follow(followerId, followeeId)`

**What it does**
* Ensures self-follow is skipped.
* Adds followee to follower’s followee set.
* Map-based set guarantees O(1) updates.

**Time Complexity O(1)**
- hash map insert.

**Space Complexity O(1)**
- adding a boolean entry.

#### `Unfollow(followerId, followeeId)`

**What it does**
* Deletes followee from follower’s followee set (if exists).

**Time Complexity O(1)**
- hash map delete.

**Space Complexity O(1)**
- no extra memory used.


### Summary Table

| Operation     | Description                  | Time Complexity    | Space Complexity |
| ------------- | ---------------------------- | ------------------ | ---------------- |
| `PostTweet`   | Append tweet with timestamp  | **O(1)**           | **O(1)**         |
| `GetNewsFeed` | K-way merge using a max heap | **O(F + K log F)** | **O(F + K)**     |
| `Follow`      | Add followee                 | **O(1)**           | **O(1)**         |
| `Unfollow`    | Remove followee              | **O(1)**           | **O(1)**         |
