import nltk
nltk.download('punkt')
from nltk.tokenize import sent_tokenize
from sentence_transformers import SentenceTransformer
model = SentenceTransformer('bert-base-nli-mean-tokens', use_auth_token='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')
from sklearn.metrics.pairwise import cosine_similarity
