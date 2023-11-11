import validateSession from "../../lib/validateSession";
import { clearCredentials } from "../../lib/credentials";
import { type FriendModel } from "../../model/friend.model";

const friendsList: HTMLUListElement | null =
  document.querySelector("#friends-list");

export default async () => {
  if (friendsList === null) {
    alert("missing element #friend-list");
    return;
  }
  const creds = await validateSession();
  if (creds === null) {
    window.alert("your session has expired");
    window.location.href = "/login/";
    clearCredentials();
    return;
  }
  while (friendsList.firstChild !== null) {
    friendsList.removeChild(friendsList.firstChild);
  }
  const friendRequest = await fetch("http://localhost:8080/friends/", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${creds.jwt}`,
    },
  });
  if (friendRequest.ok) {
    const friends = (await friendRequest.json()) as FriendModel[];
    friends
      .filter((friend) => friend.status === "friend")
      .forEach((friend) => {
        const listItem = document.createElement("li");
        listItem.classList.add(
          "my-3",
          "flex",
          "w-full",
          "items-center",
          "rounded-xl",
          "border-[1px]",
          "bg-[#4284]",
          "p-3",
          "text-black",
          "transition",
          "duration-100",
          "hover:border-red-400",
          "hover:bg-[#6f44c044]",
        );
        listItem.innerHTML = `
          <img
            class="float-left mr-3 h-16 w-16 rounded-[50%] border-[2px]"
            src="/images/user.png"
            alt=""
          />
          <div class="text-white">
            <h2>${[friend.firstname, friend.lastname].join(" ")}</h2>
          </div>
          <button
            class="ml-auto w-24 rounded-md bg-[#1877F2] text-white transition duration-150 hover:bg-[#2980b9] hover:text-black"
          >
            Friend
          </button>
        `;
        friendsList?.appendChild(listItem);
      });
  }
};
