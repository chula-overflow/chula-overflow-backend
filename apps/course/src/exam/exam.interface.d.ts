export interface ExamInterface {
  course_id: string;
  term: string; // mid | final
  semester: number; // 1, 2
  year: number; // 2022, 2021, 2020, ...
  threads:
    | {
        title: string; // if thread's length is more than 1 then generated from nlp
        description: string; // from the most upvoted thread
        problems:
          | {
              description: string;
              answer: string;
              upvoted: string;
              downvoted: string;
            }[]
          | {};
      }[]
    | {};
}
