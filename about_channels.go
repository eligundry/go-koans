package go_koans

func aboutChannels() {
  ch := make(chan string, 2)

  assert(len(ch) == 0) // channels are like buffers

  ch <- "foo" // i mean, "metaphors are like similes"

  assert(len(ch) == 1) // they can be queried for queued items

  assert(<-ch == "foo") // items can be popped out of them

  assert(len(ch) == 0) // and len() always reflects the "current" queue status
}
