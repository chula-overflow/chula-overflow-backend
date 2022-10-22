export interface ExamInterface {
  course_id: string;
  year: number; // 2022, 2021, 2020, ...
  semester: string; // S1, S2
  term: string; // mid | final
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
