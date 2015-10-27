FROM busybox
MAINTAINER Jeffrey Bean <jeffreyrobertbean@gmail.com>
ADD main main
ENTRYPOINT ["/main"]
