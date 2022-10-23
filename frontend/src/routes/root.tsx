import { WordSetList } from "../components/word-set/word-set-list";
import { useAuth } from "../context/auth";

export const Root = () => {
  const { user } = useAuth();

  return (
    <div>
      <section>Hello, {user?.name}</section>
      <WordSetList />
    </div>
  );
};
