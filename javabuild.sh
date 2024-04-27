#!/bin/sh
text="$(cat session.proto)"

package="org.baeldung.grpc"
text=${text//user_session_service/$package}
option="option java_multiple_files = true"
text=${text//option go_package=\".\"/$option}
echo "$text" > content.proto



