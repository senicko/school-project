import { useEffect, useState } from "react";
import { BookmarkIcon as BookmarkIconOutline } from "@heroicons/react/24/outline";
import { BookmarkIcon as BookmarkIconSolid } from "@heroicons/react/24/solid";
import { useUserStore } from "../../state/user";

type Joke = {
  joke: string;
};

/**
 * fetchJoke fetches a joke using an api procy to 'api.yomomma.info' api.
 * @returns fetched joke about yomomma.
 */
const fetchJoke = async (): Promise<Joke> => {
  const res = await fetch("http://localhost:3000/joke");
  return res.json();
};

/**
 * Saves the joke in user's joke collection.
 * @param joke joke that will be saved.
 */
const saveJoke = async (joke: string): Promise<void> => {
  await fetch("http://localhost:3000/users/me/jokes", {
    method: "POST",
    credentials: "include",
    body: joke,
  });
};

/**
 * JokesGenerator component fetches a joke every 10 seconds. It allows to save the joke if user is logged in.
 */
export const JokesGenerator = () => {
  const user = useUserStore((state) => state.user);
  const addJoke = useUserStore((state) => state.addJoke);

  const [joke, setJoke] = useState("");
  const [timer, setTimer] = useState(0);

  // jokeInterval
  useEffect(() => {
    fetchJoke().then(({ joke }) => {
      setJoke(joke);
      setTimer(10);
    });

    let timerInterval = setInterval(
      () => setTimer((timer) => (timer > 1 ? timer - 1 : timer)),
      1000
    );

    const jokeInterval = setInterval(() => {
      fetchJoke().then(({ joke }) => {
        setJoke(joke);
        setTimer(10);

        clearInterval(timerInterval);
        timerInterval = setInterval(
          () => setTimer((timer) => (timer > 1 ? timer - 1 : timer)),
          1000
        );
      });
    }, 10 * 1000);

    return () => {
      clearInterval(timerInterval);
      clearInterval(jokeInterval);
    };
  }, []);

  /**
   * bookmark saves the joke in user's collection.
   */
  const bookmark = async () => {
    await saveJoke(joke);
    addJoke(joke);
  };

  return (
    <div className="flex flex-col items-center justify-center gap-8">
      <div className="px-8 w-full flex justify-end">
        {user && user.jokes.some((j) => j === joke) ? (
          <BookmarkIconSolid className="h-8 w-8 " />
        ) : (
          <BookmarkIconOutline
            className="h-8 w-8 cursor-pointer"
            onClick={bookmark}
          />
        )}
      </div>
      <span className="font-medium text-5xl text-center">{joke}</span>
      <span className="italic text-3xl">next joke in {timer}</span>
    </div>
  );
};
