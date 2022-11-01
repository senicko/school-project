import { useQuery } from "@tanstack/react-query";
import { LearningSet } from "../../types";
import { Link } from "react-router-dom";

/**
 * getWordSets retrieves all word sets created by the current user.
 * @returns array of word sets
 */
const getLearningSets = async (): Promise<LearningSet[]> => {
  const res = await fetch("http://localhost:3000/word-set", {
    credentials: "include",
  });

  return await res.json();
};

/**
 * WordSetList is a component that renders all word sets created by the currently logged in user.
 */
export const LearningSetsList = () => {
  const learningSetsQuery = useQuery(["learningSets"], getLearningSets);

  if (learningSetsQuery.isLoading) return <section>Loading ...</section>;
  if (learningSetsQuery.isError)
    return <section>An error has occured.</section>;

  return (
    <section className="flex flex-col gap-2">
      <div className="flex justify-between items-center">
        <h2 className="text-xl text-gray-800 font-medium">
          Your learning sets
        </h2>
        <Link
          to="/add"
          className="flex items-center justify-center transition bg-blue-500 text-white px-3 rounded-lg hover:bg-blue-600 h-10"
        >
          Create
        </Link>
      </div>

      {learningSetsQuery.data.length == 0 ? (
        <div className="p-4 bg-gray-50 border border-gray-100 rounded-lg">
          <p className="text-gray-500 text-sm">
            It looks like you don't have any learning sets yet!
          </p>
        </div>
      ) : (
        <div>
          {learningSetsQuery.data.map(({ title }) => (
            <div className="bg-gray-50 p-4 rounded-lg border border-gray-100">
              {title}
            </div>
          ))}
        </div>
      )}
    </section>
  );
};
