package query

type Pageable interface {

	/**
	 * Returns the page to be returned.
	 *
	 * @return the page to be returned.
	 */
	GetPageNumber() int32

	/**
	 * Returns the number of items to be returned.
	 *
	 * @return the number of items of that page
	 */
	GetPageSize() int32

	/**
	 * Returns the offset to be taken according to the underlying page and page size.
	 *
	 * @return the offset to be taken
	 */

	GetOffset() int32

	/**
	* Returns the sorting parameters.
	*
	* @return
	 */
	GetSort() map[string]string

	/**
	   * Returns the current {@link Sort} or the given one if the current one is unsorted.
	   *
	   * @param sort must not be {@literal null}.
	   * @return

	  default Sort getSortOr(Sort sort) {

	  Assert.notNull(sort, "Fallback Sort must not be null!");

	  return getSort().isSorted() ? getSort() : sort;
	  }

	*/

	/**
	 * Returns the {@link Pageable} requesting the next {@link Page}.
	 *
	 * @return
	 */
	Next() Pageable

	/**
	 * Returns the previous {@link Pageable} or the first {@link Pageable} if the current one already is the first one.
	 *
	 * @return
	 */
	PreviousOrFirst() Pageable

	/**
	 * Returns the {@link Pageable} requesting the first page.
	 *
	 * @return
	 */
	First() Pageable

	/**
	 * Returns whether there's a previous {@link Pageable} we can access from the current one. Will return
	 * {@literal false} in case the current {@link Pageable} already refers to the first page.
	 *
	 * @return
	 */
	HasPrevious() bool
}

type PageableData struct {
	Page int32
	Size int32
	Sort map[string]string
}

func (p PageableData) GetPageNumber() int32 {
	return p.Page
}

func (p PageableData) GetPageSize() int32 {
	return p.Size
}

func (p PageableData) GetOffset() int32 {
	return p.Page * p.Size
}

func (p PageableData) Next() Pageable {
	panic("implement me")
}

func (p PageableData) GetSort() map[string]string {
	return p.Sort
}

func (p PageableData) PreviousOrFirst() Pageable {
	panic("implement me")
}

func (p PageableData) First() Pageable {
	p.Page = 0
	return p
}

func (p PageableData) HasPrevious() bool {
	return p.Page > 0
}
