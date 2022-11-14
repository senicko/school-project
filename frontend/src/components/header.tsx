import ProfileImage from "../assets/profile.png";
import { useUserStore } from "../state/user";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { ContextMenu } from "./context-menu";
import { Cog6ToothIcon } from "@heroicons/react/24/outline";

const logOut = async () => {
  await fetch("http://localhost:3000/users/logout", {
    credentials: "include",
  });
};

export const Header = () => {
  const navigate = useNavigate();
  const [contextMenuOpened, setContextMenuOpened] = useState(false);
  const user = useUserStore((state) => state.user);
  const setUser = useUserStore((state) => state.setUser);

  const handleLogOut = async () => {
    await logOut();
    setUser(undefined);
    navigate("/");
  };

  useEffect(() => {
    const closeContextMenu = () => setContextMenuOpened(false);
    window.addEventListener("click", closeContextMenu);
    return () => removeEventListener("click", closeContextMenu);
  }, []);

  return (
    <nav className="flex items-center justify-between gap-4">
      <div className="flex items-center gap-4">
        <img
          src={ProfileImage}
          alt="profile image"
          className="aspect-square w-[48px] rounded-full border border-gray-800"
        />
        <span className="font-medium text-gray-800">{user?.name}</span>
      </div>

      {/* <button onClick={handleLogOut}>Log Out</button> */}
      <span className="relative flex flex-col items-center justify-center">
        <Cog6ToothIcon
          className="h-6 w-6"
          onClick={(e) => {
            e.stopPropagation();
            setContextMenuOpened(true);
          }}
        />
        {contextMenuOpened && (
          <div
            className="absolute top-10 origin-center"
            onClick={(e) => e.stopPropagation()}
          >
            <ContextMenu onLogOut={handleLogOut}></ContextMenu>
          </div>
        )}
      </span>
    </nav>
  );
};
