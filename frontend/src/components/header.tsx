import ProfileImage from "../assets/profile.png";
import { useUserStore } from "../state/user";
import { useNavigate } from "react-router-dom";

const logOut = async () => {
  await fetch("http://localhost:3000/users/logout", {
    credentials: "include",
  });
};

export const Header = () => {
  const navigate = useNavigate();

  const user = useUserStore((state) => state.user);
  const setUser = useUserStore((state) => state.setUser);

  const handleLogOut = async () => {
    await logOut();
    setUser(undefined);
    navigate("/");
  };

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

      <button onClick={handleLogOut}>Log Out</button>
    </nav>
  );
};
