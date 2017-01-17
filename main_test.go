package main

import "testing"

var keyBattery = []string{"one", "1", "/test", "\\ehr", "a/somewhat/long/path"}
var valueBattery = [][]byte{
	[]byte("a value"),
	[]byte(`From a purely mathematical point of view, every frame has a budget of about 16ms (1000ms / 60 frames per second = 16.66ms per frame). However, because browsers need some time to paint the new frame to the screen, your code should finish executing in under 10ms.

In high pressure points like animations, the key is to do nothing where you can, and the absolute minimum where you can't. Whenever possible, make use of the 100ms response to pre-calculate expensive work so that you maximize your chances of hitting 60fps.

For more information, see Rendering Performance.

Idle: maximize idle time

Use idle time to complete deferred work. For example, keep pre-load data to a minimum so that your app loads fast, and use idle time to load remaining data.

Deferred work should be grouped into blocks of about 50ms. Should a user begin interacting, then the highest priority is to respond to that.

To allow for <100ms response, the app must yield control back to main thread every <50ms, such that it can execute its pixel pipeline, react to user input, and so on.

Working in 50ms blocks allows the task to finish while still ensuring instant response.

Load: deliver content under 1000ms

Load your site in under 1 second. If you don't, user attention wanders, and their perception of dealing with the task is broken.

Focus on optimizing the critical rendering path to unblock rendering.

You don't have to load everything in under 1 second to produce the perception of a complete load. Enable progressive rendering and do some work in the background. Defer non-essential loads to periods of idle time (see this Website Performance Optimization Udacity course for more information).

Summary of key RAIL metrics

To evaluate your site against RAIL metrics, use the Chrome DevTools Timeline tool to record user actions. Then check the recording times in the Timeline against these key RAIL metrics:

RAIL Step	Key Metric	User Actions
Response	Input latency (from tap to paint) < 100ms.	User taps a button (for example, opening navigation).
Animation	Each frame's work (from JS to paint) completes < 16ms.	User scrolls the page, drags a finger (to open a menu, for example), or sees an animation. When dragging, the app's response is bound to the finger position, such as pulling to refresh, or swiping a carousel. This metric applies only to the continuous phase of drags, not the start.
Idle	Main thread JS work chunked no larger than 50ms.	User isn't interacting with the page, but main thread should be available enough to handle the next user input.
Load	Page considered ready to use in 1000ms.	User loads the page and sees the critical path content.`),
	[]byte("$%&(/%&GRV%&VRT)BNT&/TB&/(\\t67b8t67 symbols!"),
	[]byte(`{"IAm": "a JSON", "key": 5647839, "the_end": true}`),
}
var contentTypeBattery []string = []string{"text/plain", "application/json", "some garbage!"}

func TestSet(t *testing.T) {
	for _, k := range keyBattery {
		for _, v := range valueBattery {
			for _, c := range contentTypeBattery {
				set(k, v, c)
			}
		}
	}
}
