export interface CourseInterface {
  course_id: string;
  course_name: string;
  course_codename: string;
  exam_ids?: string[];
}

interface course {
  _id;
  course_name;
  course_codename;
  course_id;
  exams_id;
}

interface exam {
  _id;
  course_id;
  year;
  semester;
  term;
  threads: [thread_id];
}

interface thread {
  _id;
  exam_id;
  updated_at;
  created_at;
  title;
  description;
  problems;
}
