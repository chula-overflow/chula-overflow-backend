from sentence_transformers import SentenceTransformer
model = SentenceTransformer('bert-base-nli-mean-tokens', use_auth_token='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')

def embed(sentence):
    """
    Input - (String) sentence to be embedded
    out - (numpy.ndarray) shape of (768, ) embedded vector of a sentence
    """
    return model.encode(sentence)

# print(cosine_sim(embed("which statement is true about electron and hole mobility in semiconductors?"), embed('which is true about electron and hole mobility in semiconductors?')))
