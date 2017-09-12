package main
/*
Will run the specified goroutine.
If the function returns with nil,
then all routines "owned" by that specific channel are closed.
Routines can have subroutines.
This management function is the thing called as a goroutine, callback is not.
But callback is still technically concurrent because the goroutine that surrounds it, IS concurrent.
Because object and channel will be passed to every func, bundle it into a structure.

So the rules here are:

No infinite loops. Infinite loops cannot be closed and are therefore dangerous to a program
that needs to change game states.

A loop in a goroutine looks like this:

for ShouldRun(context) {
}

This loop is as infinite as it needs to be.
All loops should use this with an AND statement as well:

for NotFinishedMoving() && ShouldRun(context) {
}

So, what parameters are passed to the function? Simple, an interface.

State is an interface. It can take any form and contain whatever data it needs.
Every state implemenets one update() method.
State can change at any time.
*/
func GameObjectManagement( this *State, callback func(*State, Context) ){
	defer remove( objects, Object ) //Remove from array if we return
	next_state:{
		ch := make(chan int)
		callback = callback( this, ch )
		terminate( ch )//Terminate all nested goroutines
		//terminate activates the SHOULD_RUN() function.
		//Or possibly, it simply takes a pointer to boolean
		//named SHOULD_QUIT. Would work just as well.
	}
	if( callback != nil ){
		goto next_state
	}		
}
