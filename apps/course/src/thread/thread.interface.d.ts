export interface CreateThreadBody {
  course_id: string;
  year: number;
  semester: string;
  term: string;

  problem: string;
  answer?: string;
}

export interface AddProblemBody {
  course_id: string;
  exam_id: string;

  year: number;
  semester: string;
  term: string;

  problem: string;
  answer: string;
}
