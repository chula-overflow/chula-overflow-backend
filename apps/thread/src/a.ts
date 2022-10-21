interface Subject {
  _id: string;
  subject_name: string;
  subject_id: string;
  subject_codename: string;
}

interface Thread {
  _id: string;
  subject_id: string;
  voted: {
    up: number;
    down: number;
  };
  problems: [
    {
      title: string; // generated from nlp
      problem: string;
      answers: string[];
    },
  ];
}
