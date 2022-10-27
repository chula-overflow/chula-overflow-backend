import nltk
nltk.download('punkt')
from nltk.tokenize import sent_tokenize

def sentence_tokenize(paragraph):
    """
    Input - (str) paragraph of question
    Output - (list of str) each sentence in list
    """
    sentences = sent_tokenize(paragraph)   
    sentences.reverse()
    return sentences

