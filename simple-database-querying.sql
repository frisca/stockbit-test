select u.id, u.username, (select us.username from "user" us where us.id=u.parent) as parentusername from "user" u;
