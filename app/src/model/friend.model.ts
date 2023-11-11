export type FriendModel = {
  firstname: string;
  lastname: string;
  status: "friend" | "sent request" | "recieved request";
};
