export interface CourseInterface {
  courseName: string;
  courseCodename: string;
  courseId: string;
  exams: {
    term: string;
    year: number;
    threads: {
      title: string; // generated from nlp
      problems: {
        description: string;
        answer: string;
        upvoted: number;
        downvoted: number;
      }[];
    }[];
  }[];
}
