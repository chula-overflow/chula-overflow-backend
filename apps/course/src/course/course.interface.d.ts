export interface CourseInterface {
  courseId: string;
  courseName: string;
  courseCodename: string;
  exams_id?: string[];
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
  term;
  semester;
  year;
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
